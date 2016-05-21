package users

import (
//"fmt"
	"../render"
	"net/http"
	"strconv"
	//str "strings"
)

func LoadTemplates() {
	//render.Load("boards", "f/boards.tmpl") // shows list of boards
	//render.Load("boardnew", "f/boardnew.tmpl")     // allows to create board
	//render.Load("boardadmin", "f/boardadmin.tmpl") // allows to set settings for existing board
	//render.Load("threads", "f/threads.tmpl") // shows whole board. for mods with extra options
	//render.Load("posts", "f/posts.tmpl")     // shows all posts in thread
	//render.Load("postedit", "f/postedit.tmpl")     // allows editing existing post
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	// yet to be implemented
	http.Error(w, "501 not implemented", 501)
}
