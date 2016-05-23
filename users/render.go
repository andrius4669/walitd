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
	page := new(pageInfo);
	page.Name = "Create Friend List";
	page.Header = "Create Friend page";
	render.Execute(w, "header", page);
	render.Execute(w, "createfriendllist", nil);
	render.Execute(w, "footer", nil);
}
func renderCreateGroupPage(w http.ResponseWriter, r *http.Request, obj *group)  {
	page := new(pageInfo);
	page.Name = "Create group";
	page.Header = "Create group page";
	render.Execute(w, "header", page);
	render.Execute(w, "creategroup", obj)
	render.Execute(w, "footer", nil);
}
func renderFriendListPage(w http.ResponseWriter, r *http.Request, friend *userAddForm, list *friendListPage)  {
	page := new(pageInfo);
	page.Name = "Friend List";
	page.Header = "Friend list page";
	render.Execute(w, "header", page);
	for i := 0; i < len(list.UsersInfo); i++ {
		render.Execute(w, "profile", list.UsersInfo[i])
	}
	render.Execute(w, "friendlist", friend)
	render.Execute(w, "removefriend", friend)
	render.Execute(w, "footer", nil);
}
func renderGroupPage(w http.ResponseWriter, r *http.Request, obj *group)  {
	page := new(pageInfo);
	page.Name = "Group";
	page.Header = "Group page";
	render.Execute(w, "header", page);
	render.Execute(w, "group", obj);
	render.Execute(w, "footer", nil);
}
func renderGroupEditPage(w http.ResponseWriter, r *http.Request,obj *group)  {
	page := new(pageInfo);
	page.Name = "Group";
	page.Header = "Group page";
	render.Execute(w, "header", page);
	render.Execute(w, "groupEdit", obj);
	render.Execute(w, "footer", nil);
}
func renderGroupsPage(w http.ResponseWriter, r *http.Request, grp *groupsPage, obj *userAddForm)  {
	page := new(pageInfo);
	page.Name = "Groups";
	page.Header = "Groups page";
	render.Execute(w, "header", page);
	for i := 0; i < len(grp.GroupsInfo); i++ {
		render.Execute(w, "group", grp.GroupsInfo[i])
	}
	for i := 0; i < len(grp.News); i++ {
		render.Execute(w, "sharedNews", grp.News[i])
	}
	render.Execute(w, "groups", obj);
	render.Execute(w, "footer", nil);
}
func renderMessagesPage(w http.ResponseWriter, r *http.Request, obj *messages, mm *messageForm)  {
	page := new(pageInfo);
	page.Name = "Messages";
	page.Header = "Messages page";
	render.Execute(w, "header", page);
	render.Execute(w, "messages", obj);
	render.Execute(w, "sendmessage", mm);
	render.Execute(w, "footer", nil);
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



