package users

import (
//	cfg "../configmgr"
//	"../files"
	"net/http"
	"../render"
	ss "../sessions"
//"path"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request, f *loginInfo)  {
	page := new(pageInfo);
	page.Name = "Login";
	page.Header = "Login page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "login", f)
	render.Execute(w, "footer", nil);
}
func renderCreateFriendListPage(w http.ResponseWriter, r *http.Request)  {
	page := new(pageInfo);
	page.Name = "Create Friend List";
	page.Header = "Create Friend page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "createfriendlist", nil);
	render.Execute(w, "footer", nil);
}
func renderCreateGroupPage(w http.ResponseWriter, r *http.Request, obj *group)  {
	page := new(pageInfo);
	page.Name = "Create group";
	page.Header = "Create group page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "creategroup", obj)
	render.Execute(w, "footer", nil);
}
func renderFriendListPage(w http.ResponseWriter, r *http.Request, friend *userAddForm, list *friendListPage, ff *friendListPage)  {
	page := new(pageInfo);
	page.Name = "Friend List";
	page.Header = "Friend list page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "text", "<h3> Users in you friend list</h3>");
	for i := 0; i < len(list.UsersInfo); i++ {
		render.Execute(w, "profile", list.UsersInfo[i])
	}
	render.Execute(w, "text", "<h3> Suggestions to add to friend list</h3>");
	for i := 0; i < len(ff.UsersInfo); i++ {
		render.Execute(w, "profile", ff.UsersInfo[i])
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
	renderMenu(w, r);
	render.Execute(w, "group", obj);
	render.Execute(w, "footer", nil);
}
func renderGroupEditPage(w http.ResponseWriter, r *http.Request,obj *group)  {
	page := new(pageInfo);
	page.Name = "Group";
	page.Header = "Group page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "groupEdit", obj);
	render.Execute(w, "footer", nil);
}
func renderGroupsPage(w http.ResponseWriter, r *http.Request, grp *groupsPage, obj *userAddForm, sug *suggests)  {
	page := new(pageInfo);
	page.Name = "Groups";
	page.Header = "Groups page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	if (len(grp.GroupsInfo) > 0) {
		render.Execute(w, "text", "<h3>You are in these groups:</h3>");
	}
	for i := 0; i < len(grp.GroupsInfo); i++ {
		grp.GroupsInfo[i].Grr = "grr";
		render.Execute(w, "group", grp.GroupsInfo[i])
	}
	render.Execute(w, "text", "<h3>We are suggesting to join these groups:</h3>");
	for i := 0; i < len(sug.Suggest); i++ {
		sug.Suggest[i].Grr = "grr";
		render.Execute(w, "group", sug.Suggest[i])
	}
	render.Execute(w, "groups", obj);
	for i := 0; i < len(grp.News); i++ {
		render.Execute(w, "sharedNews", grp.News[i])
	}
	render.Execute(w, "footer", nil);
}
func renderMessagesPage(w http.ResponseWriter, r *http.Request, obj *messages, mm *messageForm)  {
	page := new(pageInfo);
	page.Name = "Messages";
	page.Header = "Messages page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "messages", obj);
	render.Execute(w, "sendmessage", mm);
	render.Execute(w, "footer", nil);
}
func renderProfilePage(w http.ResponseWriter, r *http.Request, obj *user)  {
	page := new(pageInfo);
	page.Name = "Profile";
	page.Header = "Profile page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "profile", obj)
	render.Execute(w, "footer", nil);
}
func renderEditProfilePage(w http.ResponseWriter, r *http.Request, obj *user)  {
	page := new(pageInfo);
	page.Name = "Profile";
	page.Header = "Profile page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "profileEdit", obj)
	render.Execute(w, "footer", nil);
}
func renderRegisterPage(w http.ResponseWriter, r *http.Request, f *userForm)  {
	page := new(pageInfo);
	page.Name = "Register";
	page.Header = "Register page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "register", f)
	render.Execute(w, "footer", nil);
}
func renderMenu(w http.ResponseWriter, r *http.Request){
	ses := ss.GetUserSession(w, r);
	var ses_user_role int;
	if (ses != nil){
		uses := new(ss.UserSessionInfo);
		ss.FillUserInfo(ses, uses);
		ses_user_role = int(uses.Role);
	}
	if ses != nil {
		render.Execute(w, "menu", ses_user_role);
	}else{
		render.Execute(w, "notmenu", nil);
	}
}
func renderAdminPage(w http.ResponseWriter, r *http.Request,  f *userAddForm, ff *friendListPage)  {
	page := new(pageInfo);
	page.Name = "Register";
	page.Header = "Register page";
	render.Execute(w, "header", page);
	renderMenu(w, r);
	render.Execute(w, "admin", ff)
	render.Execute(w, "footer", nil);

}




