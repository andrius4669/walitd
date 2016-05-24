package users

import (
//	"fmt"
	"../render"
	"net/http"
	"strconv"
	str "strings"
	ss "../sessions"
//	"time"
)

func LoadTemplates() {
	render.Load("createfriendlist", "users/createFriendListLLL.tmpl");
	render.Load("creategroup", "users/creategroup.tmpl");
	render.Load("friendlist", "users/friendlist.tmpl");
	render.Load("group", "users/group.tmpl");
	render.Load("groups", "users/groups.tmpl");
	render.Load("groupEdit", "users/groupEdit.tmpl");
	render.Load("login", "users/login.tmpl");
	render.Load("messages", "users/messages.tmpl");
	render.Load("profile", "users/profile.tmpl");
	render.Load("profileEdit", "users/profileEdit.tmpl");
	render.Load("register", "users/register.tmpl");
	render.Load("header", "users/header.tmpl");
	render.Load("footer", "users/footer.tmpl");
	render.Load("removefriend", "users/removeFriend.tmpl");
	render.Load("sendmessage", "users/sendMessage.tmpl");
	render.Load("sharedNews", "users/sharedNews.tmpl");
	render.Load("menu", "menu.tmpl");
	render.Load("notmenu", "notmenu.tmpl");
	render.Load("text", "users/text.tmpl");
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

	//get session info -- also known as using magic
	ses := ss.GetUserSession(w, r);
	var ses_user_id int;
	if (ses != nil){
		uses := new(ss.UserSessionInfo);
		ss.FillUserInfo(ses, uses);
		ses_user_id = int(uses.Uid);
	} else{
		ses_user_id = -1;
	}

//	fmt.Printf("%v \n", ses);
//	fmt.Printf("%v \n", string(ses_user_id));

	rpath := r.URL.Path[pathi+1:]
//	fmt.Printf("path: %v \n sesija: %v \n if: %v \n", rpath, ses, (ses != nil));

	if r.Method == "GET" || r.Method == "HEAD" {
		p := str.IndexByte(rpath, '/')
		if p == -1 {
			// syntax is /zzz/ not /zzz in all GET/HEAD cases
			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			return
		}
		if (ses != nil || rpath[:p] == "profile"){
			if rpath == "" {
				groups := getGroupsPage(ses_user_id)
				arr := new(userAddForm) //it will be empty, in this case
				renderGroupsPage(w, r, groups, arr, getSugg(ses_user_id))
				return
			}
			i := str.IndexByte(rpath, '/')
			if i == -1 {
				// syntax is /zzz/ not /zzz in all GET/HEAD cases
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
				return
			}
			if rpath[:i] == "messages" {
				renderMessagesPage(w, r, getMessagePage(ses_user_id), new(messageForm));
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
				groups := getGroupsPage(ses_user_id);
				arr := new(userAddForm) //it will be empty, in this case
				renderGroupsPage(w, r, groups, arr, getSugg(ses_user_id))
				return
			}
			if rpath[:i] == "friendList" {
				renderFriendListPage(w, r, new(userAddForm), getFriendList());
				return
			}
			if rpath[:i] == "profile" {
				id := rpath[i+1:]
				if (id != ""){
					id, err := strconv.Atoi(id);
					if (err != nil){
						http.Redirect( w, r , "/users/profile/"+strconv.Itoa(ses_user_id), http.StatusFound);
						return;
					}
					if (ses_user_id == id){
						obj, ee := getUser(id);
						if (ee != nil){
							http.Redirect( w, r , "/users/profile/"+strconv.Itoa(ses_user_id), http.StatusFound);
						}
						renderEditProfilePage(w, r, obj);
					}else{
						obj, ee := getUser(id);
						if (ee != nil){
							http.Redirect( w, r , "/users/profile/"+strconv.Itoa(ses_user_id), http.StatusFound);
						}
						renderProfilePage(w, r, obj);
						return
					}

				} else {
					http.Redirect( w, r , "/users/", http.StatusFound);
					return
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
					gg, ee :=  getGroupPage(id);
					if ee != nil{
						http.Redirect( w, r , "/users/", http.StatusFound);
					}
					if (gg.Owner == ses_user_id){
						renderGroupPage(w, r,gg);
					} else{
						renderGroupEditPage(w, r, gg);
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
					renderLoginPage(w, r, obj);
				}
				return
			}


			http.Redirect( w, r , "/users/login/", http.StatusFound);
			return
		}
	} else if r.Method == "POST" {

		r.ParseForm()
		form := r.Form;
		i := str.IndexByte(rpath, '/')
		if (ses != nil || (rpath[:i] == "login" || rpath[:i] == "register")){
			if rpath == "" {
				obj := new(userAddForm);

				obj.Username = form["group"][0]
				act := form["act"][0];
				if (act == "join"){
					obj = joinToGroup(obj, ses_user_id);
				} else{
					obj = leaveGroup(obj, ses_user_id);
				}
				renderGroupsPage(w, r, getGroupsPage(ses_user_id), obj, getSugg(ses_user_id))
				return
			}
			if i == -1 {
				// syntax is /zzz/ not /zzz in all GET/HEAD cases
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
				return
			}
			if rpath[:i] == "messages" {
				obj := new(messageForm)
				obj.To = form["reciever"][0];
				obj.Message = form["message"][0];
				obj.Sender = strconv.Itoa(ses_user_id);
				obj = sendMessage(obj);
				renderMessagesPage(w, r, getMessagePage(ses_user_id), obj);
				return
			}
			if rpath[:i] == "login" {
				f := new(loginInfo);
				f.Username = form["username"][0];
				f.Pass = form["pass"][0];
				f.ErrorSet = false;
				login(f);

				if (f.ErrorSet){
					renderLoginPage(w, r, f);
					return;
				} else{
					loginS(w, r, f);
					http.Redirect( w, r , "/users/groups/", http.StatusFound);
				}

				return
			}
			if rpath[:i] == "logout" {
				logout(w, r);
				http.Redirect( w, r , "/users/login/", http.StatusFound);

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
					if (obj.ErrorCnt > 0){
						renderRegisterPage(w, r, obj);
					} else {
						register(obj);
						http.Redirect( w, r , "/users/login", http.StatusFound);
					}
					return
			}
			if rpath[:i] == "createfriendlist" {
				createFriendListF();
				renderCreateFriendListPage(w, r);
				return
			}
			if rpath[:i] == "creategroup" {
				obj := new(group);
				obj.Name = form["name"][0];
				obj.Description = form["desc"][0];
				var a bool;
				a = createGroup(obj, ses_user_id);
				if a{
					http.Redirect( w, r , "/users/groups/", http.StatusFound);
					return
				} else{
					renderCreateGroupPage(w, r, obj);
					return;
				}

			}
			if rpath[:i] == "groups" {
				obj := new(userAddForm);
//				fmt.Printf("%v \n", form);
				obj.Username = form["group"][0];
				act := form["act"][0];
				if (act == "join"){
					joinToGroup(obj, ses_user_id);
				} else{
					leaveGroup(obj, ses_user_id);
				}
				renderGroupsPage(w, r, getGroupsPage(ses_user_id), obj, getSugg(ses_user_id))
				return
			}
			if rpath[:i] == "friendList" {
				obj := new(userAddForm)
				obj.Username = form["user"][0];
				act := form["act"][0];
				if (act == "addFriend"){
					obj = 	addFriend(obj);
				} else{
					obj = 	removeFriend(obj);
				}
				renderFriendListPage(w, r, obj, getFriendList());
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
					arr, ee := getUser(id);
					if (ee != nil){
						http.Redirect( w, r , "/users/", http.StatusFound);
					}
//					fmt.Printf("%v \n", arr);
//					fmt.Printf("%v \n", form);

					arr.Email =form["email"][0];
					arr.FirstName =form["firstname"][0];
					arr.SecondName =form["secondname"][0];
					arr.Country =form["country"][0];
					arr.Telephone =form["telephone"][0];
					arr.City =form["city"][0];
					//TODO somehow handle birthday
//					arr.Birthday =time.Now().Format(form["birth"][0]);
					arr.Picture =form["pic"][0];
					arr.Description =form["desc"][0];
//					fmt.Printf("%v \n", arr);
					arr = validateProfileForm(arr);
					if (arr.Err > 0){
						editProfile(arr);
					}
					renderEditProfilePage(w, r, arr);
					return;
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
					obj := new(group);
					obj.Description = form["desc"][0];
					editGroup(obj, id);
					http.Redirect(w, r, r.URL.Path, http.StatusFound)
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
			http.Redirect( w, r , "/users/", http.StatusFound);
			return;
		}
		http.Error(w, "501 POST routines not yet implemented", 501)
	} else {
		http.Error(w, "501 method not implemented", 501)
	}
}
