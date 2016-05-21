package users

import (
//	"fmt"
	"../render"
	"net/http"
//	"strconv"
	str "strings"
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
// users/createfriendlist GET/POST
// users/creategroup GET/POST
// users/friendlist GET/POST
// users/group/* GET/POST
// users/groups GET
// users/login POST
// users/messages GET/POST
// users/profile/* GET/POST
// users/register GET/POST

//main page: users/groups if not loged in users/login
// * some number

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	rpath := r.URL.Path[pathi+1:]
	if r.Method == "GET" || r.Method == "HEAD" {
		if (true){ //TODO: if not logged redirect to login page
			if rpath == "" {
				renderGroupsPage(w, r)
				return
			}
			i := str.IndexByte(rpath, '/')
			if i == -1 {
				// syntax is /zzz/ not /zzz in all GET/HEAD cases
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
				return
			}
			if rpath[:i] == "messages" {
				if rpath[i+1:] == "" {
					// Display list of news
					renderMessagesPage(w, r)
				} else {
				}
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
		} else{
			i := str.IndexByte(rpath, '/')
			if rpath[:i] == "register" {
				if rpath[i+1:] == "" {
					// Display list of news
					renderRegisterPage(w, r)
				}
				return
			}
			// else render login page
			renderLoginPage(w, r);
		}
	} else if r.Method == "POST" {
		// TODO(andrius) decide & implement posting
		http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}
