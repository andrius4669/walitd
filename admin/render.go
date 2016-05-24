package admin


import (
"../render"
//"fmt"
"net/http"
//"time"
"../dbacc"
)

//func Execute(w io.Writer, name string, data interface{})
func renderAdminPage(w http.ResponseWriter, r *http.Request) {
//page := new(blank)
render.Execute(w, "headerA", nil);
render.Execute(w, "ArticlesFunctions", nil);
render.Execute(w, "footer", nil);
}

func renderEditPage(w http.ResponseWriter, r *http.Request){
	//page := new(articlesContent)
	render.Execute(w, "headerA", nil)
	render.Execute(w, "choose", nil)
	render.Execute(w, "footer", nil)
}
func renderEditPageFull(w http.ResponseWriter, r *http.Request, name string){
	page := new(articlesList)
	db := dbacc.OpenSQL()
	defer db.Close()
	getArticle(db, page, name)
	render.Execute(w, "headerA", nil)
	//render.Execute(w, "choose", nil)
	render.Execute(w, "edit", page)
	render.Execute(w, "footer", nil)
}