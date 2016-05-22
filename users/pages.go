package users


import (
	"time"
)

type loginPage struct{

}
type registerPage struct{

}

type profilePage struct{
	UserInfo user;
}
type user struct {
	Userid int;
	Email string;
	EmailErr string
	FirstName string;
	FirstNameErr string
	SecondName string;
	SecondNameErr string
	Username string;
	Pass string;
	PassErr string;
	Role int;
	RoleN string;
	Birthday time.Time;
	City string;
	Country string;
	Telephone string;
	Gender int;
	GenderN string;
	Description string;
	Created time.Time;
	Updated time.Time;
	Picture string;
	PictureCreatedPictureCreated time.Time;
	Err int;
}
func (u user) p(){
	u.Err = u.Err + 1;
}
type sharedNews struct{
	Sharedid int;
	Userid int;
	Newsid int;
	Groupname string;
	Desc string;
	Created time.Time;
}


type groupsPage struct{
	GroupsInfo [] group;
	News [] sharedNews;
}
type createGroupPage struct {
	GroupInfo group;
}
type group struct {
	GroupId int;
	Name string;
	NameErr string;
	Created time.Time;
	GroupType int;
	Updated time.Time;
	Description string;
	ErrCnt int;
}

func getGroup(id int) *group  {
	g := new(group);
	g.Name = "grupe";
	g.Description ="zasias";
	return g;
}
type friendListPage struct {
	UsersInfo [] user;
}
type createFriendList struct {
	UsersInfo [] user;
}
type groupPage struct {
	GroupInfo group;
}
type messagesPage struct  {
	MessagesInfo [] message;
}
type message struct {
	MessageId int;
	Sender string;
	Reciever string;
	SenderId int;
	RecieverId int;
	Text string;
	Created time.Time;
}


type userForm struct {
	Email string;
	EmailErr string;
	FirstName string;
	FirstNameErr string;
	SecondName string;
	SecondNameErr string;
	Username string;
	UsernameErr string;
	Pass string;
	PassErr string;
	City string;
	CityErr string;
	Country string;
	CountryErr string;
	Gender int;
	GenderErr string;
	ErrorCnt int;

}
func (m userForm) p() {
	m.ErrorCnt = m.ErrorCnt + 1;
}
func makeErrorMessage(m string) string{
	return "<p class='error'>" + m + "</p>";
}

func (m userForm) IsMale() bool{
	if (m.Gender == 0){
		return true;
	}
	return false;
}
func (m userForm) IsFemale() bool {
	if (m.Gender == 1){
		return true;
	}
	return false;
}
func (m userForm) IsDunno() bool {
	if (m.Gender == 2){
		return true;
	}
	return false;
}

type pageInfo struct{
	Name string;
	Header string;
}
type loginInfo struct{
	Username string;
	Pass string;
	Error string;
	ErrorSet bool;
}

func getUser(id int) *user  {
	obj := new(user);
	obj.Userid = 1;
	obj.Email = "mail@mail.com";
	obj.FirstName = "pirmas ardas";
	obj.SecondName = " antrs names";
	obj.Username = "userio prisijungimas";
	obj.RoleN = "rol4";
	obj.City = "kaunas";
	obj.Country = "LL";
	obj.Telephone = "584848";
	obj.GenderN = "LL";
	obj.Description = "desc";
	obj.Picture = "adresas";
	return obj;
}