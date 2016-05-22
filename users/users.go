package users

import (
//	"fmt"
	"../render"
	"net/http"
	"strconv"
	str "strings"
//	"time"
)

func LoadTemplates() {
	render.Load("createfriendlist", "users/createfriendlist.tmpl");
	render.Load("creategroup", "users/creategroup.tmpl");
	render.Load("friendlist", "users/friendlist.tmpl");
	render.Load("group", "users/group.tmpl");
	render.Load("groupEdit", "users/groupEdit.tmpl");
	render.Load("login", "users/login.tmpl");
	render.Load("messages", "users/messages.tmpl");
	render.Load("profile", "users/profile.tmpl");
	render.Load("profileEdit", "users/profileEdit.tmpl");
	render.Load("register", "users/register.tmpl");
	render.Load("header", "users/header.tmpl");
	render.Load("footer", "users/footer.tmpl");
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

//	fmt.Println(r.Form) // print information on server side.

//	fmt.Printf("%v \n", r);
//	fmt.Printf("%v \n", r.Method);
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
				renderMessagesPage(w, r);
				return
			}
			if rpath[:i] == "createfriendlist" {
				renderCreateFriendListPage(w, r);
				return
			}
			if rpath[:i] == "creategroup" {
				obj := new(group);
				renderCreateGroupPage(w, r, obj);
				return
			}
			if rpath[:i] == "groups" {
				renderGroupsPage(w, r);
				return
			}
			if rpath[:i] == "friendList" {
				renderFriendListPage(w, r);
				return
			}
			if rpath[:i] == "profile" {
				id := rpath[i+1:]
				if (id != ""){
					id, err := strconv.Atoi(id);
					if (err != nil){
						http.Redirect( w, r , "/users/", http.StatusFound);
						return;
					}
					//TODO: check if user's own profile
					if (true){
						obj := getUser(id);
						renderEditProfilePage(w, r, obj);
					}else{
						//TODO get obj from database;
						obj := getUser(id);
						renderProfilePage(w, r, obj);
					}

				} else {
					http.Redirect( w, r , "/users/", http.StatusFound);
				}
				return
			}
			if rpath[:i] == "group" {
				id := rpath[i+1:]
				if (id != ""){
					id, err := strconv.Atoi(id);
					if (err != nil){
						http.Redirect( w, r , "/users/", http.StatusFound);
						return;
					}
					if (true){
						renderGroupPage(w, r, id);
					} else{
						renderGroupEditPage(w, r, id);
					}

				} else {
					http.Redirect( w, r , "/users/", http.StatusFound);
				}
				return
			}


			if rpath == "" {
				http.Redirect( w, r , "/users/", http.StatusFound);
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
			if i == -1 {
				// syntax is /zzz/ not /zzz in all GET/HEAD cases
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
				return
			}
			if rpath[:i] == "register" {
				if rpath[i+1:] == "" {
					renderRegisterPage(w, r, new(userForm))
				}
				return
			}
			if rpath[:i] == "login" {
				if rpath[i+1:] == "" {
					obj := new(loginInfo);
					renderLoginPage(w, r,obj);
				}
				return
			}


			http.Redirect( w, r , "/users/groups", http.StatusFound);
		}
	} else if r.Method == "POST" {

		r.ParseForm()
		form := r.Form;
		i := str.IndexByte(rpath, '/')
		if (true && (rpath[:i] == "login" || rpath[:i] == "register")){ //TODO: if not logged redirect to login page if not post form login or register
			if rpath == "" {
				//TODO: handle post request
				renderGroupsPage(w, r)
				return
			}
			if i == -1 {
				// syntax is /zzz/ not /zzz in all GET/HEAD cases
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
				return
			}
			if rpath[:i] == "messages" {
				//TODO: handle post request
				renderMessagesPage(w, r);
				return
			}
			if rpath[:i] == "login" {
				//TODO: handle post request.
				f := new(loginInfo);
				f.Username = form["username"][0];
				f.Pass = form["pass"][0];
				f.ErrorSet = false;
				obj := validateLoginForm(f);
				if (obj.ErrorSet){
					renderLoginPage(w, r, f);
				} else{
					http.Redirect( w, r , "/users/groups", http.StatusFound);
				}

				return
			}
			if rpath[:i] == "register" {
					arr := new(userForm)
					arr.Pass = form["pass"][0];
					arr.Country = form["country"][0];
					arr.Gender , _ = strconv.Atoi(form["gender"][0]);
					arr.FirstName = form["firstname"][0];
					arr.SecondName = form["lastname"][0];
					arr.Username = form["username"][0];
					arr.Email = form["email"][0];
					arr.City = form["town"][0];
					obj := validateRegisterForm(arr);
					if (obj.ErrorCnt > 0 || true){ //TODO: check why error counter doesnt work
						//TODO: save form to database
						renderRegisterPage(w, r, obj);
					} else {
						http.Redirect( w, r , "/users/login", http.StatusFound);
					}
					return
			}
			if rpath[:i] == "createfriendlist" {
				//TODO: handle post request
				renderCreateFriendListPage(w, r);
				return
			}
			if rpath[:i] == "creategroup" {
				//TODO: handle post request
				obj := new(group);
				renderCreateGroupPage(w, r, obj);
				return
			}
			if rpath[:i] == "groups" {
				//TODO: handle post request
				renderGroupsPage(w, r);
				return
			}
			if rpath[:i] == "friendList" {
				//TODO: handle post request
				renderFriendListPage(w, r);
				return
			}
			if rpath[:i] == "profile" {
				id := rpath[i+1:]
				if (id != ""){
					id, err := strconv.Atoi(id);
					if (err != nil){
						http.Redirect( w, r , "/users/", http.StatusFound);
						return;
					}
					//TODO: handle post request
					arr := getUser(id);
					arr.Email =form["email"][0];
					arr.FirstName =form["firstname"][0];
					arr.SecondName =form["secondname"][0];
					arr.Country =form["country"][0];
					arr.Telephone =form["telephone"][0];
					arr.City =form["city"][0];
//					arr.Birthday =time.Now().Format(form["birth"][0]);
					arr.Picture =form["pic"][0];
					arr.Description =form["desc"][0];
					renderEditProfilePage(w, r, arr);
				} else {
					http.Redirect( w, r , "/users/", http.StatusFound);
				}
				return
			}
			if rpath[:i] == "group" {
				id := rpath[i+1:]
				if (id != ""){
					id, err := strconv.Atoi(id);
					if (err != nil){
						http.Redirect( w, r , "/users/", http.StatusFound);
						return;
					}
					//TODO: handle post request
					renderGroupPage(w, r, id);
				} else {
					http.Redirect( w, r , "/users/", http.StatusFound);
				}
				return
			}


			if rpath == "" {
				http.Redirect( w, r , "/users/", http.StatusFound);
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
					renderRegisterPage(w, r, new(userForm))
				}
				return
			}
			// else render login page
			obj := new(loginInfo);
			renderLoginPage(w, r, obj);
		}
		http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}
