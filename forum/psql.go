package forum

// sql related stuff
// currently only github.com/lib/pq lib usage

import (
	"../users"
	"database/sql"
	"encoding/json"
	//"fmt"
	"strconv"
	"time"
)

func queryBoardList(db *sql.DB, p *frontPage) {
	rows, err := db.Query("SELECT bname, topic, description FROM forum.boards")
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
	err := db.QueryRow("SELECT boardid, topic, description, attributes FROM forum.boards WHERE bname=$1", board).Scan(&bid, &p.Topic, &p.Description, &attributesjson)
	if err == sql.ErrNoRows {
		//fmt.Printf("fail @ initial select: no rows\n")
		return false
	}
	panicErr(err)
	p.Board = board

	type attributes struct {
		PageLimit      *uint32
		ThreadsPerPage *uint32
		AllowNewThread *bool
		AllowFiles     *bool
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
		//fmt.Printf("fail @ page check\n")
		return false
	}

	rows, err := db.Query("SELECT threadid, bump FROM forum.threads WHERE boardid=$1 ORDER BY bump DESC LIMIT $2 OFFSET $3", bid, tpp, (page-1)*tpp)
	panicErr(err)
	for rows.Next() {
		var t threadInfo
		rows.Scan(&t.ID, &t.Bump)
		p.Threads = append(p.Threads, t)
	}

	// if no threads and no limit, only show existing threads
	if len(p.Threads) == 0 && page != 1 && pagelimit == 0 {
		//fmt.Printf("fail @ threads num check: len=%d page=%d pl=%d\n", len(p.Threads), page, pagelimit)
		return false
	}

	var allthreads uint32
	err = db.QueryRow("SELECT COUNT(*) FROM forum.threads WHERE boardid=$1", bid).Scan(&allthreads)
	panicErr(err)

	var cp uint32
	p.Pages = append(p.Pages, true)
	for cp = 1; cp < (allthreads+tpp-1)/tpp; cp++ {
		p.Pages = append(p.Pages, true)
	}
	if pagelimit != 0 {
		for ; cp < pagelimit; cp++ {
			p.Pages = append(p.Pages, false)
		}
	}

	p.CurrentPage = page

	for i := range p.Threads {
		// get stuff from OP
		var uid sql.NullInt64
		err = db.QueryRow("SELECT title, userid, pname, trip, email FROM forum.posts WHERE boardid=$1 AND postid=$2", bid, p.Threads[i].ID).Scan(&p.Threads[i].Title, &uid, &p.Threads[i].OP.Name, &p.Threads[i].OP.Trip, &p.Threads[i].OP.Email)
		panicErr(err)
		if uid.Valid {
			p.Threads[i].OP.User = users.GetUserInfo(uint32(uid.Int64))
		}
		// get info about last post
		err = db.QueryRow("SELECT postid, userid, pname, trip, email, postdate FROM forum.posts WHERE boardid=$1 AND threadid=$2 ORDER BY postdate DESC LIMIT 1", bid, p.Threads[i].ID).Scan(&p.Threads[i].LastID, &uid, &p.Threads[i].Last.Name, &p.Threads[i].Last.Trip, &p.Threads[i].Last.Email, &p.Threads[i].LastDate)
		panicErr(err)
		if uid.Valid {
			p.Threads[i].Last.User = users.GetUserInfo(uint32(uid.Int64))
		}
		// get reply count
		var numrepl uint32
		err = db.QueryRow("SELECT COUNT(*) FROM forum.posts WHERE boardid=$1 AND threadid=$2", bid, p.Threads[i].ID).Scan(&numrepl)
		panicErr(err)
		p.Threads[i].Replies = numrepl - 1
	}

	p.Mod = mod

	return true
}

func queryPostFiles(p *postContent, db *sql.DB) {
}

func queryThread(db *sql.DB, p *threadPage, board, thread string, mod bool) bool {
	// sanity checks first
	if !validBoardName(board) {
		return false
	}

	var tid uint32
	var err error
	// no casting for multiple return values. no nested funcs too. fuck you golang.
	// I hope this shit will atleast get inlined.
	intcst := func(i uint64, e error) (uint32, error) {
		return uint32(i), e
	}
	tid, err = intcst(strconv.ParseUint(thread, 10, 32))
	if err != nil {
		return false
	}

	var attributesjson []byte
	var bid uint32
	err = db.QueryRow("SELECT boardid, topic, description, attributes FROM forum.boards WHERE bname=$1", board).Scan(&bid, &p.Topic, &p.Description, &attributesjson)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	p.Board = board

	err = db.QueryRow("SELECT threadid FROM forum.threads WHERE boardid=$1 AND threadid=$2", bid, tid).Scan(&p.ID)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)

	rows, err := db.Query("SELECT postid, userid, pname, trip, email, title, postdate, message FROM forum.posts WHERE boardid=$1 AND threadid=$2 ORDER BY postdate ASC", bid, tid)
	panicErr(err)

	p.refMap = make(map[uint32]int)
	for rows.Next() {
		var pc postContent
		var uid sql.NullInt64
		err = rows.Scan(&pc.PostID, &uid, &pc.UserIdent.Name, &pc.UserIdent.Trip, &pc.UserIdent.Email, &pc.Title, &pc.Date, &pc.Message)
		panicErr(err)
		if uid.Valid {
			pc.UserIdent.User = users.GetUserInfo(uint32(uid.Int64))
		}

		if pc.PostID != tid {
			p.Replies = append(p.Replies, pc)
			p.refMap[pc.PostID] = len(p.Replies)
		} else {
			p.OP = pc
			p.refMap[pc.PostID] = 0
		}
	}

	queryPostFiles(&p.OP, db)
	for i := range p.Replies {
		queryPostFiles(&p.Replies[i], db)
	}

	formatPost(p, &p.OP, db)
	for i := range p.Replies {
		formatPost(p, &p.Replies[i], db)
	}

	p.Mod = mod

	return true
}

func sqlGetBoard(db *sql.DB, board string) (uint32, bool) {
	if !validBoardName(board) {
		return 0, false
	}
	if reservedBoardName(board) {
		return 0, false
	}
	var bid uint32
	err := db.QueryRow("SELECT boardid FROM forum.boards WHERE bname=$1", board).Scan(&bid)
	if err == sql.ErrNoRows {
		return 0, false
	}
	panicErr(err)
	return bid, true
}

func sqlValidateBoard(db *sql.DB, board string) bool {
	_, ok := sqlGetBoard(db, board)
	return ok
}

func sqlValidatePost(db *sql.DB, board string, post uint32, thread *uint32) bool {
	bid, ok := sqlGetBoard(db, board)
	if !ok {
		return false
	}
	var tid sql.NullInt64
	err := db.QueryRow("SELECT threadid FROM forum.posts WHERE boardid=$1 AND postid=$2", bid, post).Scan(&tid)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	if tid.Valid {
		*thread = uint32(tid.Int64)
	} else {
		*thread = post
	}
	return true
}

func validateInputBoard(db *sql.DB, d *boardData) bool {
	if !validBoardName(d.Board) {
		return false
	}
	if reservedBoardName(d.Board) {
		return false
	}
	if d.Topic == "" {
		return false
	}
	return true
}

func validateInputThread(db *sql.DB, d *postMessage) bool {
	bid, ok := sqlGetBoard(db, d.Board)
	if !ok {
		return false
	}
	d.boardid = bid
	return true
}

func validateInputPost(db *sql.DB, d *postData) bool {
	bid, ok := sqlGetBoard(db, d.Board)
	if !ok {
		return false
	}
	d.boardid = bid
	var tid uint32
	err := db.QueryRow("SELECT threadid FROM forum.threads WHERE boardid=$1 AND threadid=$2", bid, d.ThreadID).Scan(&tid)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	return true
}

func sqlStoreBoard(db *sql.DB, d *boardData) bool {
	stmt, err := db.Prepare("INSERT INTO forum.boards (bname, topic, description, attributes) VALUES ($1, $2, $3, $4)")
	panicErr(err)
	type attributes struct {
		PageLimit      uint32 `json:,omitempty`
		ThreadsPerPage uint32 `json:,omitempty`
		AllowNewThread bool
		AllowFiles     bool
	}
	var attr attributes
	attr.PageLimit = d.PageLimit
	attr.ThreadsPerPage = d.ThreadsPerPage
	attr.AllowNewThread = d.AllowNewThread
	attr.AllowFiles = d.AllowFiles
	encattr, _ := json.Marshal(attr)
	//fmt.Printf("encattr=%s\n", encattr)
	_, err = stmt.Exec(d.Board, d.Topic, d.Description, encattr) // TODO
	panicErr(err)
	return true
}

func sqlStoreThread(db *sql.DB, d *postMessage) bool {
	// allocate new id
	var tid uint32
	err := db.QueryRow("UPDATE forum.boards SET lastid = lastid+1 WHERE boardid=$1 RETURNING lastid", d.boardid).Scan(&tid)
	panicErr(err)
	// insert into threads list
	stmt, err := db.Prepare("INSERT INTO forum.threads (boardid, threadid, bump) VALUES ($1, $2, $3)")
	panicErr(err)
	nowtime := time.Now()
	_, err = stmt.Exec(d.boardid, tid, nowtime)
	panicErr(err)
	// insert into posts list
	stmt, err = db.Prepare("INSERT INTO forum.posts (boardid, postid, threadid, userid, pname, trip, email, title, postdate, message) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	panicErr(err)
	_, err = stmt.Exec(d.boardid, tid, tid, d.UserID, d.PName, d.Trip, d.Email, d.Title, nowtime, d.Message)
	panicErr(err)
	// done
	return true
}

func sqlStorePost(db *sql.DB, d *postData) bool {
	// allocate new id
	var pid uint32
	err := db.QueryRow("UPDATE forum.boards SET lastid = lastid+1 WHERE boardid=$1 RETURNING lastid", d.boardid).Scan(&pid)
	panicErr(err)
	nowtime := time.Now()
	if d.bump {
		// bump thread
		stmt, err := db.Prepare("UPDATE forum.threads SET bump=$1 WHERE boardid=$2 AND threadid=$3")
		panicErr(err)
		_, err = stmt.Exec(nowtime, d.boardid, d.ThreadID)
		panicErr(err)
	}
	// insert into posts list
	stmt, err := db.Prepare("INSERT INTO forum.posts (boardid, postid, threadid, userid, pname, trip, email, title, postdate, message) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	panicErr(err)
	_, err = stmt.Exec(d.boardid, pid, d.ThreadID, d.UserID, d.PName, d.Trip, d.Email, d.Title, nowtime, d.Message)
	panicErr(err)
	// done
	return true
}
