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
