package news

import (
	"fmt"
	"../render"
	"net/http"
	str "strings"
	"strings"
	"strconv"
	//"database/sql"
	"time"
	"../dbacc"
)

func LoadTemplates() {
	//render.Load("boards", "f/boards.tmpl") // shows list of boards
	//render.Load("boardnew", "f/boardnew.tmpl")     // allows to create board
	//render.Load("boardadmin", "f/boardadmin.tmpl") // allows to set settings for existing board
	//render.Load("threads", "f/threads.tmpl") // shows whole board. for mods with extra options
	//render.Load("posts", "f/posts.tmpl")     // shows all posts in thread
	render.Load("article", "articles/article.tmpl")     // allows editing existing post
	render.Load("list", "articles/list.tmpl")
	render.Load("createArticle", "articles/createArticle.tmpl")
	render.Load("searchArticle", "articles/searchArticle.tmpl")
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	rpath := r.URL.Path[pathi+1:]
	if r.Method == "GET" || r.Method == "HEAD" {
		if rpath == "" {
			// Display list of Articles
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
				//renderArticles(w, r, rpath[i+1:])
				renderArticlesList(w, r)
			} else {
				// display moderation page for specific board. possibly will check for admin
				//renderArticles(w, r, rpath[i+1:])s[:len(s)-len(suffix)]
				//fmt.Printf("%v \n", rpath[:len(rpath) - (len(rpath) - 1)]);
				//fmt.Printf("%v \n", rpath[:len(rpath)]);
				//fmt.Printf("%v \n", rpath[i+1:len(rpath) - 1]);

				temp, err := strconv.Atoi(rpath[i+1:len(rpath) - 1])
				fmt.Printf("%v \n", err);
				renderArticles(w, r, temp)
			}
			return
		}
		if rpath[:i] == "createArticle" {
			renderArticleCreation(w, r)
			return
		}
		if rpath[:i] == "searchArticle" {
			renderArticleSearch(w, r)
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


		http.NotFound(w, r)
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		form := r.Form;
		i := str.IndexByte(rpath, '/')
		//fmt.Printf("%v \n", rpath[:i]);
		//fmt.Printf("%v \n", i);

		if rpath[:i] == "articles" {
			vote := form["vote"][0];
			fmt.Printf("%v \n", vote);
			temp, err := strconv.Atoi(rpath[i+1:len(rpath) - 1])
			fmt.Printf("%v \n", err);
			renderArticles(w, r, temp)
			return
		}
		if rpath[:i] == "createArticle" {
			arr := new(articlesList)
			arr.Name = form["name"][0];
			arr.Description = form["description"][0];
			arr.Article = form["article"][0];
			arr.Category = form["category"][0];
			arr.Author = 1
			tim := time.Now()
			tim.Format("2006-01-02 15:04:05")
			arr.UploadDate = tim.Format("2006-01-02")
			tags := form["tags"][0];
			noCaseTags := strings.ToLower(tags)
			slices := strings.Split(noCaseTags, ",")
			//fmt.Printf("%v \n", name);
			//fmt.Printf("%v \n", description);
			//fmt.Printf("%v \n", article);
			fmt.Printf("%v \n", tags);
			fmt.Printf("%v \n", slices);
			fmt.Printf("%v \n", slices[1]);
			db := dbacc.OpenSQL()
			defer db.Close()
			createArticle(db, arr)
			http.Redirect( w, r , "/news/", http.StatusFound);
			return
		}
		if rpath[:i] == "searchArticle" {
			name := form["name"][0];
			author := form["author"][0];
			tags := form["tags"][0];
			noCaseTags := strings.ToLower(tags)
			slices := strings.Split(noCaseTags, ",")
			fmt.Printf("%v \n", name);
			fmt.Printf("%v \n", author);
			fmt.Printf("%v \n", tags);
			fmt.Printf("%v \n", slices);
			fmt.Printf("%v \n", slices[1]);

			renderArticlesList(w, r)
			return
		}
		//http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}