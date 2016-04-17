package main_rdr

import (
	cfg "../configmgr"
	"../files"
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "HEAD" {
		// redirect links without trailing / to proper places
		switch r.URL.Path[1:] {
		case "users":
			fallthrough
		case "forum":
			fallthrough
		case "news":
			fallthrough
		case "poll":
			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			return
		}
		// serve root
		if r.URL.Path == "/" {
			fn, ok := cfg.GetOption("main.mainpage")
			if ok {
				files.ServeFileOr404(w, r, fn)
			} else {
				fmt.Fprintf(w, "Nothing to see there")
			}
			return
		}
		// else, 404
		http.Error(w, fmt.Sprintf("404 page not found: %s", r.URL.Path), 404)
		//http.NotFound(w, r)
	} else {
		http.Error(w, "501 not implemented", 501)
	}
}
