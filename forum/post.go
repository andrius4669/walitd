package forum

import (
	"../dbacc"
	"fmt"
	"net/http"
	sc "strconv"
)

// new board creation
func handleNewBoard(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	d := new(boardData)
	bn, _ := r.Form["board"]
	if len(bn) > 0 {
		d.Board = bn[0]
	}
	tpc, _ := r.Form["topic"]
	if len(bn) > 0 {
		d.Topic = tpc[0]
	}
	dsc, _ := r.Form["desc"]
	if len(bn) > 0 {
		d.Description = dsc[0]
	}
	pl, _ := r.Form["pagelimit"]
	if len(pl) > 0 && pl[0] != "" {
		i, err := sc.Atoi(pl[0])
		if err == nil && i > 0 {
			d.PageLimit = uint32(i)
		}
	}
	tpp, _ := r.Form["threadsperpage"]
	if len(tpp) > 0 && tpp[0] != "" {
		i, err := sc.Atoi(tpp[0])
		if err == nil && i > 0 {
			d.ThreadsPerPage = uint32(i)
		}
	}
	yes, _ := r.Form["allownewthread"]
	if len(yes) > 0 && yes[0] == "yes" {
		d.AllowNewThread = true
	}
	yes, _ = r.Form["allowfiles"]
	if len(yes) > 0 && yes[0] == "yes" {
		d.AllowFiles = true
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	if !validateInputBoard(db, d) {
		fmt.Fprintf(w, "bad data")
		return
	}
	if sqlStoreBoard(db, d) {
		fmt.Fprintf(w, "board made")
	} else {
		fmt.Fprintf(w, "failed to make board")
	}
}

// new thread or post creation
func handlePost(w http.ResponseWriter, r *http.Request) {
	//err = r.ParseMultipartForm(1 << 20)
	//d := new(boardData)
}
