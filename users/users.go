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
// users/createfriendlist
// users/creategroup
// users/friendlist
// users/group/*
// users/groups
// users/login
// users/messages
// users/profile/*
// users/register

//main page: users/groups if not loged in users/login
// * some number
