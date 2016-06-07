package news

import (
	"../render"
	"net/http"
	"../dbacc"
)

func renderArticlesList(w http.ResponseWriter, r *http.Request) {
	page := new(ArticlesFrontPage)
	db := dbacc.OpenSQL()
	defer db.Close()
	queryArticlesList(db, page)
	render.Execute(w, "list", page)
}

func renderArticles(w http.ResponseWriter, r *http.Request, id int) {
	page := new(articlesList)
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
	page := new(articlesContent)
	render.Execute(w, "createArticle", page)
}

func renderArticleSearch(w http.ResponseWriter, r *http.Request){
	page := new(articlesContent)
	render.Execute(w, "searchArticle", page)
}