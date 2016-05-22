package users

import (
//	cfg "../configmgr"
//	"../files"
	"net/http"
	"../render"
//"path"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request, f *loginInfo)  {
	page := new(pageInfo);
	page.Name = "Login";
	page.Header = "Login page";
	render.Execute(w, "header", page);
	render.Execute(w, "login", f)
	render.Execute(w, "footer", nil);
}
func renderCreateFriendListPage(w http.ResponseWriter, r *http.Request)  {
	http.Error(w, "Register page not implemented", 501);
}
func renderCreateGroupPage(w http.ResponseWriter, r *http.Request, obj *group)  {
	page := new(pageInfo);
	page.Name = "Create group";
	page.Header = "Create group page";
	render.Execute(w, "header", page);
	render.Execute(w, "creategroup", obj)
	render.Execute(w, "footer", nil);
}
func renderFriendListPage(w http.ResponseWriter, r *http.Request)  {
	http.Error(w, "Friend list page not implemented", 501);
}
func renderGroupPage(w http.ResponseWriter, r *http.Request, id int)  {
	http.Error(w, "Group  page not implemented", 501);
}
func renderGroupEditPage(w http.ResponseWriter, r *http.Request, id int)  {
	http.Error(w, "Group  page not implemented", 501);
}
func renderGroupsPage(w http.ResponseWriter, r *http.Request)  {
	http.Error(w, "Groups page not implemented", 501);
}
func renderMessagesPage(w http.ResponseWriter, r *http.Request)  {
	http.Error(w, "Messages page not implemented", 501);
}
func renderProfilePage(w http.ResponseWriter, r *http.Request, obj *user)  {
	page := new(pageInfo);
	page.Name = "Profile";
	page.Header = "Profile page";
	render.Execute(w, "header", page);
	render.Execute(w, "profile", obj)
	render.Execute(w, "footer", nil);
}
func renderEditProfilePage(w http.ResponseWriter, r *http.Request, obj *user)  {
	page := new(pageInfo);
	page.Name = "Profile";
	page.Header = "Profile page";
	render.Execute(w, "header", page);

	render.Execute(w, "profileEdit", obj)
	render.Execute(w, "footer", nil);
}
func renderRegisterPage(w http.ResponseWriter, r *http.Request, f *userForm)  {
	page := new(pageInfo);
	page.Name = "Register";
	page.Header = "Register page";
	render.Execute(w, "header", page);
	render.Execute(w, "register", f)
	render.Execute(w, "footer", nil);
}



