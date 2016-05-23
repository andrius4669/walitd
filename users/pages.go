package users


import (
	"time"
	"../dbacc"
	"strconv"
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
	UsernameErr string;
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
	PictureCreated time.Time;
	Err int;
}
func (u user) p(){
	u.Err = u.Err + 1;
}
type sharedNews struct{
	Sharedid int;
	Userid int;
	Newsid int;
	Title string;
	UserName string;
	Groupname string;
	Desc string;
	Created time.Time;
	Link string;
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
	Owner int;
	OwnerName string;
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
type messages struct {
	Recieved []message;
	Sent []message;
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
func (m *userForm) p() {
	m.ErrorCnt = m.ErrorCnt + 1;
}
func makeErrorMessage(m string) string{
	return "<p class='error'>" + m + "</p>";
}

func (m *userForm) IsMale() bool{
	if (m.Gender == 0){
		return true;
	}
	return false;
}
func (m *userForm) IsFemale() bool {
	if (m.Gender == 1){
		return true;
	}
	return false;
}
func (m *userForm) IsDunno() bool {
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

func getUser(id int) (*user, error)  {
	db := dbacc.OpenSQL();
	defer db.Close();
	idd := strconv.Itoa(id);
	//TODO get dynamic info
	obj := new(user);
	err := queryGetUser(db, obj, idd);
	return obj, err;
}

type messageForm struct {
	Sender string;
	SenderErr string;
	To string;
	Message string;
}
type userAddForm struct {
	Username string;
	UsernameErr string;
}
func getGroupsPage() *groupsPage{
	//TODO: return groups page info
	return new(groupsPage);
}
func getMessagePage() *messages{
	//TODO: return Messages
	return new(messages);
}
func getFriendList() *friendListPage {
	//TODO: return friend list
	return new(friendListPage);
}
func getGroupPage(id int) (*group, error){
	db := dbacc.OpenSQL();
	defer db.Close();
	gg := new(group);
	err := queryGetGroup(db, gg, id);
	return gg, err;
}
func joinToGroup(gr *userAddForm) *userAddForm{
	//TODO: join group
	return gr;
}
func leaveGroup(gr *userAddForm) *userAddForm {
	//TODO: leave group
	return gr;
}
func sendMessage(m *messageForm) *messageForm {
	//TODO send message
	return m;
}
func register(r *userForm) {
	db := dbacc.OpenSQL();
	defer db.Close();
	queryAddUser(db, r);
}
func createFriendListF()  {
	//TODO create friend list
}
func createGroup(g *group) ( bool){
	//TODO check username and create group
	if (true){
		return true;
	} else{
		return false;
	}
}
func addFriend(o *userAddForm) *userAddForm{
	//TODO add friend
	return o;
}
func removeFriend(o *userAddForm) *userAddForm{
	//TODO remove friend
	return o;
}
func editProfile(obj *user){
	db := dbacc.OpenSQL();
	defer db.Close();
	queryUpdateUser(db, obj);
}
func editGroup(obj *group, id int){
	db := dbacc.OpenSQL();
	defer db.Close();
	queryUpdateGroup(db, obj, id);
}