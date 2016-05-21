package users

import (
//	"fmt"
	"../render"
//	"net/http"
//	"strconv"
	//str "strings"
)

func LoadTemplates() {
	render.Load("createfriendlist", "users/createfriendlist.tmpl")
	render.Load("creategroup", "users/creategroup.tmpl")
	render.Load("friendlist", "users/friendlist.tmpl")
	render.Load("group", "users/group.tmpl")
	render.Load("login", "users/login.tmpl")
	render.Load("messages", "users/messages.tmpl")
	render.Load("profile", "users/profile.tmpl")
	render.Load("register", "users/register.tmpl")
}

//func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
//	rpath := r.URL.Path[pathi+1:]
//	if r.Method == "GET" || r.Method == "HEAD" {
//		if rpath == "" {
//			// display list of boards
////			renderBoardList(w, r)
//			//TODO: show main page
//			return
//		}
//		i := str.IndexByte(rpath, '/')
//		if i == -1 {
//			// syntax is /zzz/ not /zzz in all GET/HEAD cases
//			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
//			return
//		}
//
//	} else if r.Method == "POST" {
//		// TODO(andrius) decide & implement posting
//		http.Error(w, "501 POST routines not yet implemented", 501)
//	} else {
//		http.Error(w, "501 method not implemented", 501)
//	}
//}
//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()  // parse arguments, you have to call this by yourself
//	fmt.Println(r.Form)  // print form information in server side
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
//}
//
//func main() {
//	http.HandleFunc("/", sayhelloName) // set router
//	err := http.ListenAndServe(":9090", nil) // set listen port
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}
