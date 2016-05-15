package forum

import (
	"../render"
	"fmt"
	"net/http"
	//"time"
)

// TODO(andrius)
//func Execute(w io.Writer, name string, data interface{})
func renderBoardList(w http.ResponseWriter, r *http.Request) {
	page := new(frontPage)
	page.Boards = append(page.Boards, boardInfo{Board: "test", Topic: "testinfo", Description: "test desc"})
	page.Boards = append(page.Boards, boardInfo{Board: "test2", Topic: "testinfo2", Description: "test desc2"})
	render.Execute(w, "boards", page)
}

func renderBoardListModPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 board list mod page not implemented", 501)
}

func renderBoardModPage(w http.ResponseWriter, r *http.Request, board string) {
	http.Error(w, fmt.Sprintf("501 board %s mod page not implemented", board), 501)
}

func renderBoardPage(w http.ResponseWriter, r *http.Request, board string, pid uint32, mod bool) {
	page := new(boardPage)
	page.Mod = mod
	page.Board = board
	page.Topic = "test topic"
	page.Description = "some description describing this test board"
	for i := uint32(1); i <= 5; i++ {
		var t threadInfo
		t.ID = i
		t.Title = "test title"
		t.Replies = 100 + i
		page.Threads = append(page.Threads, t)
	}
	for i := 1; i <= 5; i++ {
		page.Pages = append(page.Pages, true)
	}
	for i := 1; i <= 5; i++ {
		page.Pages = append(page.Pages, false)
	}
	page.CurrentPage = pid
	render.Execute(w, "threads", page)
	//http.Error(w, fmt.Sprintf("501 board %s page %d (mod: %t) not implemented", board, page, mod), 501)
}

func renderThread(w http.ResponseWriter, r *http.Request, board string, thread string, mod bool) {
	page := new(threadPage)
	page.Mod = mod
	page.Board = board
	page.Topic = "test topic"
	page.Description = "some description describing this test board"
	page.Thread.ID = 123
	var p postContent
	p.PostID = 123
	p.Title = "lol this is title"
	p.Message = "lol this is message"
	p.FMessage = "lol this is message"
	p.References = append(p.References, 124)
	p.References = append(p.References, 125)
	p.Files = append(p.Files, fileContent{Name: "123test.png", Original: "test original.png", Thumb: "/forum/"+board+"/static/testthumb.jpg"})
	page.Thread.OP = p
	p.PostID = 124
	p.UserIdent.Name = "wandalizorours"
	p.UserIdent.Trip = "!aksa6df54a1"
	p.References = nil
	p.References = append(p.References, 125)
	p.Files = nil
	page.Thread.Replies = append(page.Thread.Replies, p)
	p.PostID = 125
	p.UserIdent.Email = "sage"
	p.References = nil
	p.Files = append(p.Files, fileContent{Name: "123test.png", Original: "test original.png", Thumb: "/forum/"+board+"/static/testthumb.jpg"})
	p.Files = append(p.Files, fileContent{Name: "/deleted", Original: "test original.png", Thumb: "/forum/"+board+"/static/testthumb.jpg"})
	p.Files = append(p.Files, fileContent{Name: "123test.png", Original: "test original.png", Thumb: "/forum/"+board+"/static/testthumb.jpg"})
	page.Thread.Replies = append(page.Thread.Replies, p)
	p.PostID = 126
	p.UserIdent.Name = ""
	page.Thread.Replies = append(page.Thread.Replies, p)
	p.PostID = 127
	p.UserIdent.Email = ""
	page.Thread.Replies = append(page.Thread.Replies, p)
	render.Execute(w, "posts", page)
	//http.Error(w, fmt.Sprintf("501 board %s thread %s (mod: %t) not implemented", board, thread, mod), 501)
}
