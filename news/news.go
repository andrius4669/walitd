package news

import (
	"fmt"
	"../render"
	"net/http"
	str "strings"
	"strings"
	"strconv"
	"time"
	"../dbacc"
	ss "../sessions"
)

func LoadTemplates() {
	render.Load("article", "articles/article.tmpl")     // allows editing existing post
	render.Load("list", "articles/list.tmpl")
	render.Load("createArticle", "articles/createArticle.tmpl")
	render.Load("searchArticle", "articles/searchArticle.tmpl")
	render.Load("searchResult", "articles/searchResult.tmpl")
	render.Load("footer", "articles/footer.tmpl")
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	rpath := r.URL.Path[pathi+1:]
	if r.Method == "GET" || r.Method == "HEAD" {
		if rpath == "" {
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
				renderArticlesList(w, r)
			} else {
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


		i = str.IndexByte(rpath, '/')

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
		if rpath[:i] == "articles" {
			vote := form["vote"][0];
			temp, err := strconv.Atoi(rpath[i+1:len(rpath) - 1])
			fmt.Printf("%v err: \n", err);

			ses := ss.GetUserSession(w, r);
			var ses_user_id int;
			if (ses != nil){
				uses := new(ss.UserSessionInfo);
				ss.FillUserInfo(ses, uses);
				ses_user_id = int(uses.Uid);

				db := dbacc.OpenSQL()
				defer db.Close()
				voteForArticle(db, temp, ses_user_id, vote)
			}


			renderArticles(w, r, temp)
			return
		}
		if rpath[:i] == "createArticle" {
			arr := new(articlesList)
			arr.Name = form["name"][0];
			arr.Description = form["description"][0];
			arr.Article = form["article"][0];
			arr.Category = form["category"][0];

			ses := ss.GetUserSession(w, r);
			var ses_user_id int;
			if (ses != nil) {
				uses := new(ss.UserSessionInfo);
				ss.FillUserInfo(ses, uses);
				ses_user_id = int(uses.Uid);
			}

			arr.Author = ses_user_id
			tim := time.Now()
			tim.Format("2006-01-02 15:04:05")
			arr.UploadDate = tim.Format("2006-01-02")
			tags := form["tags"][0];
			noCaseTags := strings.ToLower(tags)
			slices := strings.Split(noCaseTags, ",")

			db := dbacc.OpenSQL()
			var article_id int
			defer db.Close()
			createArticle(db, arr)
			getArticleID(db, arr, slices, article_id)
			createTags(db, slices)
			http.Redirect( w, r , "/news/", http.StatusFound);
			return
		}
		if rpath[:i] == "searchArticle" {
			tags := form["tags"][0];
			noCaseTags := strings.ToLower(tags)
			slices := strings.Split(noCaseTags, ",")

			renderSearchResult(w, r, slices)
			return
		}

		//http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}