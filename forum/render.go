package forum

import (
	"../render"
	"fmt"
	"net/http"
	//"time"
	"../dbacc"
	ss "../sessions"
)

// TODO(andrius)
//func Execute(w io.Writer, name string, data interface{})
func renderBoardList(w http.ResponseWriter, r *http.Request) {
	page := new(frontPage)

	db := dbacc.OpenSQL()
	defer db.Close()

	queryBoardList(db, page)

	render.Execute(w, "boards", page)
}

func renderBoardListModPage(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "501 board list mod page not implemented", 501)
	s := ss.GetUserSession(w, r)
	if s == nil {
		http.Error(w, "401 unauthorized: not logged in", 401)
		return
	}
	usi := new(ss.UserSessionInfo)
	ss.FillUserInfo(s, usi)
	if usi.Role < 2 {
		http.Error(w, "401 unauthorized: privilege too low", 401)
		return
	}
	page := new(frontPage)
	page.Mod = true

	db := dbacc.OpenSQL()
	defer db.Close()

	queryBoardList(db, page)

	render.Execute(w, "boards", page)
}

func renderBoardModPage(w http.ResponseWriter, r *http.Request, board string) {
	http.Error(w, fmt.Sprintf("501 board %s mod page not implemented", board), 501)
}

func renderBoardPage(w http.ResponseWriter, r *http.Request, board string, pid uint32, mod bool) {
	if mod {
		s := ss.GetUserSession(w, r)
		if s == nil {
			http.Error(w, "401 unauthorized: not logged in", 401)
			return
		}
		usi := new(ss.UserSessionInfo)
		ss.FillUserInfo(s, usi)
		if usi.Role < 2 {
			http.Error(w, "401 unauthorized: privilege too low", 401)
			return
		}
	}
	page := new(boardPage)
	page.Mod = mod
	db := dbacc.OpenSQL()
	defer db.Close()
	if !queryBoard(db, page, board, pid, mod) {
		http.NotFound(w, r)
		return
	}
	render.Execute(w, "threads", page)
}

func renderThread(w http.ResponseWriter, r *http.Request, board string, thread string, mod bool) {
	if mod {
		s := ss.GetUserSession(w, r)
		if s == nil {
			http.Error(w, "401 unauthorized: not logged in", 401)
			return
		}
		usi := new(ss.UserSessionInfo)
		ss.FillUserInfo(s, usi)
		if usi.Role < 2 {
			http.Error(w, "401 unauthorized: privilege too low", 401)
			return
		}
	}
	page := new(threadPage)
	page.Mod = mod
	db := dbacc.OpenSQL()
	defer db.Close()
	if !queryThread(db, page, board, thread, mod) {
		http.NotFound(w, r)
		return
	}
	render.Execute(w, "posts", page)
}
