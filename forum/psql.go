package forum

// sql related stuff
// currently only github.com/lib/pq lib usage

import (
	//"fmt"
	"encoding/json"
	"database/sql"
	"strconv"
	"../users"
)

func queryBoardList(db *sql.DB, p *frontPage) {
	rows, err := db.Query("SELECT bname, topic, description FROM boards")
	panicErr(err)

	for rows.Next() {
		var b boardInfo
		rows.Scan(&b.Board, &b.Topic, &b.Description)
		p.Boards = append(p.Boards, b)
	}
}

func queryBoard(db *sql.DB, p *boardPage, board string, page uint32, mod bool) bool {
	// sanity checks first
	if !validBoardName(board) {
		return false
	}

	var attributesjson []byte
	var bid uint32
	err := db.QueryRow("SELECT boardid, topic, description, attributes FROM boards WHERE bname=$1", board).Scan(&bid, &p.Topic, &p.Description, &attributesjson)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	p.Board = board

	type attributes struct {
		PageLimit *uint32
		ThreadsPerPage *uint32
		AllowNewThread *bool
		AllowFiles *bool
	}
	var attr attributes
	json.Unmarshal(attributesjson, &attr) // fail shouldn't happen

	var tpp uint32
	if attr.ThreadsPerPage != nil && *attr.ThreadsPerPage > 0 {
		tpp = *attr.ThreadsPerPage
	} else {
		tpp = 10 // default
	}

	var pagelimit uint32
	if attr.PageLimit != nil {
		pagelimit = *attr.ThreadsPerPage
	} else {
		pagelimit = 0 // default -- no limit
	}

	if page == 0 || (pagelimit != 0 && page > pagelimit) {
		return false
	}

	rows, err := db.Query("SELECT threadid, bump FROM threads WHERE boardid=$1 ORDER BY bump DESC LIMIT $2 OFFSET $3", bid, tpp, (page - 1) * tpp)
	panicErr(err)
	for rows.Next() {
		var t threadInfo
		rows.Scan(&t.ID, &t.Bump)
		p.Threads = append(p.Threads, t)
	}

	// if no threads and no limit, only show existing threads
	if len(p.Threads) == 0 && page != 1 && pagelimit == 0 {
		return false
	}

	var allthreads uint32
	err = db.QueryRow("SELECT COUNT(*) FROM threads WHERE boardid=$1", bid).Scan(&allthreads)
	panicErr(err)

	var cp uint32
	for cp = 0; cp < (allthreads + tpp - 1) / tpp; cp++ {
		p.Pages = append(p.Pages, true)
	}
	if pagelimit != 0 {
		for ; cp < pagelimit; cp++ {
			p.Pages = append(p.Pages, false)
		}
	} else if len(p.Pages) == 0 {
		p.Pages = append(p.Pages, false)
	}

	p.CurrentPage = page

	for i := range p.Threads {
		// get stuff from OP
		var uid sql.NullInt64
		err = db.QueryRow("SELECT title, userid, pname, trip, email FROM posts WHERE boardid=$1 AND postid=$2", bid, p.Threads[i].ID).Scan(&p.Threads[i].Title, &uid, &p.Threads[i].OP.Name, &p.Threads[i].OP.Trip, &p.Threads[i].OP.Email)
		panicErr(err)
		if uid.Valid {
			p.Threads[i].OP.User = users.GetUserInfo(uint32(uid.Int64))
		}
		// get info on last post
		err = db.QueryRow("SELECT postid, userid, pname, trip, email, postdate FROM posts WHERE boardid=$1 AND threadid=$2 ORDER BY postdate DESC LIMIT 1", bid, p.Threads[i].ID).Scan(&p.Threads[i].LastID, &uid, &p.Threads[i].Last.Name, &p.Threads[i].Last.Trip, &p.Threads[i].Last.Email, &p.Threads[i].LastDate)
		panicErr(err)
		if uid.Valid {
			p.Threads[i].Last.User = users.GetUserInfo(uint32(uid.Int64))
		}
	}

	return true
}

func queryThread(db *sql.DB, board, thread string) bool {
	// sanity checks first
	if !validBoardName(board) {
		return false
	}

	//var tid uint32
	var err error
	// no casting for multiple return values. no nested funcs too. fuck you golang.
	// I hope this shit will atleast get inlined.
	intcst := func(i uint64, e error) (uint32, error) {
		return uint32(i), e
	}
	_, err = intcst(strconv.ParseUint(thread, 10, 32))
	if err != nil {
		return false
	}

	// TODO(andrius)
	return false
}
