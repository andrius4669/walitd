package news

import (
	"../render"
	"net/http"
	str "strings"
	//"strconv"
)

func LoadTemplates() {
	//render.Load("boards", "f/boards.tmpl") // shows list of boards
	//render.Load("boardnew", "f/boardnew.tmpl")     // allows to create board
	//render.Load("boardadmin", "f/boardadmin.tmpl") // allows to set settings for existing board
	//render.Load("threads", "f/threads.tmpl") // shows whole board. for mods with extra options
	//render.Load("posts", "f/posts.tmpl")     // shows all posts in thread
	//render.Load("postedit", "f/postedit.tmpl")     // allows editing existing post
	render.Load("list", "articles/list.tmpl")
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	rpath := r.URL.Path[pathi+1:]
	if r.Method == "GET" || r.Method == "HEAD" {
		if rpath == "" {
			// display list of boards
			//renderBoardList(w, r)
			renderArticlesList(w, r)
			return
		}
		i := str.IndexByte(rpath, '/')
		if i == -1 {
			// syntax is /zzz/ not /zzz in all GET/HEAD cases
			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			return
		}
		if rpath[:i] == "articles" {
			if rpath[i+1:] == "" {
				// Display list of news
				renderArticlesList(w, r)
			} else {
				// display moderation page for specific board. possibly will check for admin
				//renderArticles(w, r, rpath[i+1:])
			}
			return
		}
		/*if rpath[:i] == "static" {
			if rpath[i+1:] != "" {
				serveStatic(w, r, rpath[i+1:])
			} else {
				// be lazy there :^)
				http.Redirect(w, r, "../", http.StatusFound)
			}
			return
		}
		*/
		/*board := rpath[:i]
		rpath = rpath[i+1:]
		mod := false
		    if rpath == "mod" {
			// append /
			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			return
		} */
		/* if len(rpath) >= 4 && rpath[:4] == "mod/" {
			mod = true
			rpath = rpath[4:]
		}
		*/
		if rpath == "" {
			// render first page
			//renderBoardPage(w, r, board, 1, mod)
			return
		}
		i = str.IndexByte(rpath, '/')

		/*
		if i < 0 {
			n, err := strconv.ParseUint(rpath, 10, 32)
			if err == nil {
				// render nth page
	//			renderBoardPage(w, r, board, uint32(n), mod)
			} else {
				// append /
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			}
			return
		}
		*/
		//group := rpath[:i]
		//rpath = rpath[i+1:]

		if rpath == "" {
			// be lazy there :^)
			http.Redirect(w, r, "../", http.StatusFound)
			return
		}

		/* if group == "thumb" {
			// server thumbnail
			serveThumb(w, r, board, rpath)
			return
		} else if group == "static" {
			// serve static file
			serveBoardStatic(w, r, board, rpath)
			return
		} */

		/*
		if i = str.IndexByte(rpath, '/'); i != -1 {
			// ignore anything including and past /
			rpath = rpath[:i]
		}
		*/

		/*if group == "thread" {
			// render specific thread
			renderThread(w, r, board, rpath, mod)
			return
		} else if group == "src" {
			// serve source file
			serveSrc(w, r, board, rpath)
			return
		}*/
		http.NotFound(w, r)
		return
	} else if r.Method == "POST" {
		// TODO(andrius) decide & implement posting
		http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}