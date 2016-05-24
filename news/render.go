package news

import (
	"../render"
	//"fmt"
	"net/http"
	//"time"
	"../dbacc"
)

//func Execute(w io.Writer, name string, data interface{})
func renderArticlesList(w http.ResponseWriter, r *http.Request) {
	page := new(ArticlesFrontPage)
	//page.Boards = append(page.Boards, articlesList{Author: "test", Category: "testinfo", Description: "test desc", FullArticleLink: "TESTY"})
	//page.Boards = append(page.Boards, articlesList{Author: "test2", Category: "testinfo2", Description: "test desc2", FullArticleLink: "TESTY2"})

	db := dbacc.OpenSQL()
	defer db.Close()
	queryArticlesList(db, page)
	render.Execute(w, "list", page)
}

func renderArticles(w http.ResponseWriter, r *http.Request, id int) {
	page := new(articlesList)
	//page.UserIdent.UserName = "Algirdas-Lukas Narbutas"
	//page.UserIdent.Date
	//page.Date
	db := dbacc.OpenSQL()
	defer db.Close()
	queryArticle(db, page, id)

	render.Execute(w, "article", page)
}

func renderSearchResult(w http.ResponseWriter, r *http.Request, search []string) {
	page := new(ArticlesFrontPage)

	db := dbacc.OpenSQL()
	defer db.Close()
	queryArticlesSearchList(db, page, search)

	render.Execute(w, "searchArticle", page)
	render.Execute(w, "searchResult", page)
	render.Execute(w, "footer", page)
}

func renderArticleCreation(w http.ResponseWriter, r *http.Request){
	page := new(articlesContent) //random struc
	//db := dbacc.OpenSQL()
	//defer db.Close()
	//createArticle(db, page)
	render.Execute(w, "createArticle", page)
}

func renderArticleSearch(w http.ResponseWriter, r *http.Request){
	page := new(articlesContent) //random struc
	render.Execute(w, "searchArticle", page)
}