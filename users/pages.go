package users


import (
	"time"
)

type loginPage struct{

}
type registerPage struct{

}

type profilePage struct{
	userInfo user;
}
type user struct {
	userid uint32;
	email string;
	firstName string;
	secondName string;
	role uint32;
	birthDate time.Time;
	city string;
	country string;
	telephone string;
	gender uint32;
	description string;
	created time.Time;
	updated time.Time;
	picture string;
	pictureCreated time.Time;
}
type sharedNews struct{
	sharedid uint32;
	userid uint32;
	newsid uint32;
	groupname string;
	desc string;
	created time.Time;
}


type groupsPage struct{
	groupsInfo [] group;
	news [] sharedNews;
}
type createGroupPage struct {
	groupInfo group;
}
type group struct {
	groupId uint32;
	name string;
	created time.Time;
	groupType uint32;
	updated time.Time;
	description string;
}
type friendListPage struct {
	usersInfo [] user;
}
type createFriendList struct {
	usersInfo [] user;
}
type groupPage struct {
	groupInfo group;
}
type messagesPage struct  {
	messagesInfo [] message;
}
type message struct {
	messageId uint32;
	sender string;
	reciever string;
	senderId uint32;
	recieverId uint32;
	text string;
	created time.Time;
}
