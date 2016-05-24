package admin

import (

"../render"
"net/http"
str "strings"
//"strings"
//"strconv"
	"../dbacc"
)

func LoadTemplates() {
//render.Load("boards", "f/boards.tmpl") // shows list of boards
//render.Load("boardnew", "f/boardnew.tmpl")     // allows to create board
//render.Load("boardadmin", "f/boardadmin.tmpl") // allows to set settings for existing board
//render.Load("threads", "f/threads.tmpl") // shows whole board. for mods with extra options
	render.Load("headerA", "admin/headerA.tmpl")     // shows all posts in thread
	render.Load("ArticlesFunctions", "admin/ArticlesFunctions.tmpl")
	render.Load("footer", "admin/footer.tmpl")
	render.Load("edit", "admin/editArticle.tmpl")
	render.Load("choose", "admin/chooseArticle.tmpl")
}
//-- No check wheather user is actualy an admin
func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
rpath := r.URL.Path[pathi+1:]
if r.Method == "GET" || r.Method == "HEAD" {
	if rpath == "" {
	// Display admin page
		renderAdminPage(w, r, )
	return
	}
	i := str.IndexByte(rpath, '/')
	if i == -1 {
	// syntax is /zzz/ not /zzz in all GET/HEAD cases
		http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
		return
	}
	if rpath[:i] == "edit" {
		renderEditPage(w, r);
		return
	}


if rpath == "" {
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

	if rpath[:i] == "edit" {
		edit := form["name"][0];
		//temp, err := strconv.Atoi(rpath[i+1:len(rpath) - 1])
		//fmt.Printf("%v \n", err);
		renderEditPageFull(w, r, edit)


		//b := new(articlesList)
		//b.Name = form["name"][0]
		//b.Description = form["description"][0]
		//b.Article = form["article"][0]
		//db := dbacc.OpenSQL()
		//defer db.Close()
		//updateArticle(db, b)

		return
	}
	if rpath[:i] == "update" {
		b := new(articlesList)
		b.Name = form["name"][0]
		b.Description = form["description"][0]
		b.Article = form["article"][0]
		//temp, err := strconv.Atoi(rpath[i+1:len(rpath) - 1])
		//fmt.Printf("%v \n", err);
		db := dbacc.OpenSQL()
		defer db.Close()
		updateArticle(db, b)
		renderEditPageFull(w, r, b.Name)
		return
	}

	delete := form["delete"][0];

	db := dbacc.OpenSQL()
	defer db.Close()
	queryDeleteArticle(db, delete);
	renderAdminPage(w, r)
	if delete != ""	{

	}

} else {
http.Error(w, "501 method not implemented", 501)
}
}