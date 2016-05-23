package admin


import (
"../render"
//"fmt"
"net/http"
//"time"
)

//func Execute(w io.Writer, name string, data interface{})
func renderAdminPage(w http.ResponseWriter, r *http.Request) {
//page := new(blank)
render.Execute(w, "headerA", nil);
render.Execute(w, "ArticlesFunctions", nil);
render.Execute(w, "footer", nil);
}

