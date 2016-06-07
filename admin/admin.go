package admin

import (
"fmt"
"../render"
"net/http"
str "strings"

	"../dbacc"
)

func LoadTemplates() {
	render.Load("headerA", "admin/headerA.tmpl")     // shows all posts in thread
	render.Load("ArticlesFunctions", "admin/ArticlesFunctions.tmpl")
	render.Load("footer", "admin/footer.tmpl")
	render.Load("edit", "admin/editArticle.tmpl")
	render.Load("choose", "admin/chooseArticle.tmpl")
}

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
http.Redirect(w, r, "../", http.StatusFound)
return
}

http.NotFound(w, r)
return
} else if r.Method == "POST" {
	r.ParseForm()
	form := r.Form;
	i := str.IndexByte(rpath, '/')


	delete := form["delete"][0];
	fmt.Printf("%v DELETE \n", delete);
	if delete != "" {
		db := dbacc.OpenSQL()
		defer db.Close()
		queryDeleteArticle(db, delete);
		renderAdminPage(w, r)
	}


	if i == -1 {
		// syntax is /zzz/ not /zzz in all GET/HEAD cases
		http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
		return
	}

	if rpath[:i] == "edit" {
		edit := form["name"][0];

		fmt.Printf("%v RPATH \n", rpath[i+1:]);
			if rpath[i+1:] == "update" {
				b := new(articlesList)
				b.Name = form["name"][0]
				b.Description = form["description"][0]
				b.Article = form["article"][0]
				db := dbacc.OpenSQL()
				defer db.Close()
				id := new(int)
				b.ID = getArticleID(db, b, id)
				updateArticle(db, b)
			}

		renderEditPageFull(w, r, edit)

		return
	} else if rpath[:i] == "update" {
		b := new(articlesList)
		b.Name = form["name"][0]
		b.Description = form["description"][0]
		b.Article = form["article"][0]
		db := dbacc.OpenSQL()
		defer db.Close()
		updateArticle(db, b)
		renderEditPageFull(w, r, b.Name)
		return
	} else {

		delete := form["delete"][0];
		fmt.Printf("%v DELETE \n", delete);
		db := dbacc.OpenSQL()
		defer db.Close()
		queryDeleteArticle(db, delete);
		renderAdminPage(w, r)
		if delete != "" {

		}
	}

} else {
http.Error(w, "501 method not implemented", 501)
}
}