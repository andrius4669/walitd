package admin


import (
"../render"
//"fmt"
"net/http"
//"time"
)

//func Execute(w io.Writer, name string, data interface{})
func renderAdminPage(w http.ResponseWriter, r *http.Request, f *deleteValidate) {
//page := new(blank)
render.Execute(w, "headerA", nil);
render.Execute(w, "ArticlesFunctions", f);
render.Execute(w, "footer", nil);
}

