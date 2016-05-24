package users
import (
//"fmt"
//	"encoding/json"
	"database/sql"
//	"strconv"
)
func queryGetUser(db *sql.DB, u *user, id string) error{
	var email, city, country, phone, desc, picture sql.NullString;
	err := db.QueryRow("SELECT userid, username, email, firstname, lastname, role, city, country, telephone, gender, description, created, updated, photo, photocreated FROM users where userid=$1;", id).
	Scan(&u.Userid, &u.Username, &email , &u.FirstName, &u.SecondName, &u.Role, &city, &country, &phone, &u.Gender, &desc, &u.Created, &u.Updated, &picture, &u.PictureCreated);
	if (email.Valid){
		u.Email = email.String;
	}
	if (city.Valid){
		u.City = city.String;
	}
	if (country.Valid){
		u.Country = country.String;
	}
	if (phone.Valid){
		u.Telephone = phone.String;
	}
	if (desc.Valid){
		u.Description = desc.String;
	}
	if (picture.Valid){
		u.Picture = picture.String;
	}
	if (u.Gender == 0){
		u.GenderN ="Male"
	}
	if (u.Gender == 1){
		u.GenderN ="Female"
	}
	if (u.Gender == 2){
		u.GenderN ="Uncertain"
	}
	if (u.Role == 1){
		u.RoleN ="User"
	}
	if (u.Role == 2){
		u.RoleN ="Admin"
	}
	if (u.Role == 3) {
		u.RoleN = "Super Admin"
	}
//	panicErr(err);
	return err;
}
func queryGetUserByUsername(db *sql.DB,u *user, username string) error {
	var email, city, country, phone, desc, picture sql.NullString;
	err := db.QueryRow("SELECT userid, username, email, firstname, lastname, role, city, country, telephone, gender, description, created, updated, photo, photocreated FROM users where username=$1;", username).
	Scan(&u.Userid, &u.Username, &email , &u.FirstName, &u.SecondName, &u.Role, &city, &country, &phone, &u.Gender, &desc, &u.Created, &u.Updated, &picture, &u.PictureCreated);
	if (email.Valid){
		u.Email = email.String;
	}
	if (city.Valid){
		u.City = city.String;
	}
	if (country.Valid){
		u.Country = country.String;
	}
	if (phone.Valid){
		u.Telephone = phone.String;
	}
	if (desc.Valid){
		u.Description = desc.String;
	}
	if (picture.Valid){
		u.Picture = picture.String;
	}
	if (u.Gender == 0){
		u.GenderN ="Male"
	}
	if (u.Gender == 1){
		u.GenderN ="Female"
	}
	if (u.Gender == 2){
		u.GenderN ="Uncertain"
	}
	if (u.Role == 1){
		u.RoleN ="User"
	}
	if (u.Role == 2){
		u.RoleN ="Admin"
	}
	if (u.Role == 3){
		u.RoleN ="Super Admin"
	}
//	panicErr(err);
	return err;
}
func queryGetMessages(db *sql.DB, m *messages, id int){
	rows, err := db.Query("Select sender, username, message, messages.created from messages left join users on userid=reciever where sender=$1;", id);
	panicErr(err);
	for rows.Next() {
		var t message
		rows.Scan(&t.Sender, &t.Reciever, &t.Text, &t.Created);
		m.Sent = append(m.Sent, t);
	}
	rows, err = db.Query("Select username, reciever, message, messages.created from messages left join users on userid=sender where reciever=$1;", id);
	panicErr(err);
	for rows.Next() {
		var t message
		rows.Scan(&t.Sender, &t.Reciever, &t.Text, &t.Created);
		m.Recieved = append(m.Recieved, t);
	}
}
func queryGetGroup(db *sql.DB, g *group, id int) error{
	var desc sql.NullString;
	g.GroupId = id;
	err := db.QueryRow("Select groupid, name, description, created, updated from groups where groupid=$1 and grouptype=1;", id).Scan(&g.GroupId, &g.Name, &desc, &g.Created, &g.Updated);
	if (desc.Valid){
		g.Description = desc.String;
	}
	if err != nil {
		return err;
	}
	pp := db.QueryRow("select users.userid, users.username from usergroup left join users on users.userid=usergroup.userid where level=1 and groupid=$1", id).Scan(&g.Owner, &g.OwnerName);
	return pp;
}
func queryGetGroupByName(db *sql.DB, g *group, id string) error{
	var desc sql.NullString;
	err := db.QueryRow("Select groupid, name, description, created, updated from groups where name=$1 and grouptype=1;", id).Scan(&g.GroupId, &g.Name, &desc, &g.Created, &g.Updated);
	if err != nil {
		return err;
	}
	if (desc.Valid){
		g.Description = desc.String;
	}

	pp := db.QueryRow("select users.userid, users.username from usergroup left join users on users.userid=usergroup.userid where level=1 and groupid=$1", g.GroupId).Scan(&g.Owner, &g.OwnerName);
	return pp;
}
func queryCreateGroup(db *sql.DB, g *group, ownerid int)  {
	db.Query("insert into groups (name, description, created, grouptype, updated) values ($1, $2, now(), 1, now())", g.Name, g.Description);
//	panicErr(err);
	queryGetGroupByName(db, g, g.Name);
	db.Query("insert into usergroup (groupid, userid, level, created) values($1, $2, 1, now())", g.GroupId, ownerid);
}
func queryAddMessage(db *sql.DB, m *messageForm, uid int)  {
	_, err := db.Query("insert into messages (sender, reciever, message, created) values($1, $2, $3, now())", m.Sender, uid, m.Message);
	panicErr(err);
}
func queryAddUser(db *sql.DB, u *userForm){
	_, err := db.Query("insert into users (username, password, firstname, lastname, role, gender, created, updated, photocreated, active, email) values($1, $2, $3, $4, $5, $6, now(), now(), now(), 1, $7);", u.Username, u.Pass, u.FirstName, u.SecondName, 1, u.Gender, u.Email);
	panicErr(err);
}
func queryUpdateUser(db *sql.DB, u *user){
//	fmt.Printf("%v \n", u);
	_, err := db.Query("update users set email=$1, firstname=$2, lastname=$3, country=$4, telephone=$5, city=$6, description=$8, updated=now() where userid=$9", u.Email, u.FirstName, u.SecondName, u.Country, u.Telephone, u.City, u.Description, u.Userid);
	panicErr(err);
	if (u.Picture != ""){
		_, err := db.Query("update users set photo=$2, photocreated=now() where userid=$1", u.Userid, u.Picture);
		panicErr(err);
	}
}
func queryUpdateGroup(db *sql.DB, g *group, id int)  {
	_, err := db.Query("update groups set description=$2, updated=now() where groupid=$1", id, g.Description);
	panicErr(err);
}
func queryGetGroupList(db *sql.DB,g *groupsPage, userid int) {
	rows, err := db.Query("Select groups.groupid from usergroup left join groups on usergroup.groupid=groups.groupid where usergroup.userid=$1 and groups.grouptype=1", userid);
	panicErr(err);
	for rows.Next() {
		var t int;
		rows.Scan(&t);
		gg := new(group);
		queryGetGroup(db, gg, t);
		g.GroupsInfo = append(g.GroupsInfo, *gg);
	}
}
func queryCheckLogin(db *sql.DB, l *loginInfo) error{
	var username sql.NullString;
	err := db.QueryRow("Select username from users where username=$1 and password=$2", l.Username, l.Pass).Scan(&username);
//	panicErr(err);
	return err;
}
func queryLeaveGroup(db *sql.DB, groupid int, id int){
	_, err:= db.Query("delete from usergroup where groupid=$1 and userid=$2", groupid, id);
	panicErr(err);
}
func queryDestroyGroup(db *sql.DB, groupid int){
	db.Query("delete from usergroup where groupid=$1", groupid);
	db.Query("delete from groups where groupid=$1", groupid);
}
func queryCheckUserGroup(db *sql.DB, gg *group, userid int) error{
	err := db.QueryRow("select groupid from usergroup where userid=$1 and groupid=$2", userid, gg.GroupId).Scan(&gg.GroupId);
	return err;
}
func queryJoinGroup(db *sql.DB, gid int, uid int) {
	db.Query("insert into usergroup (groupid, userid, level, created) values($1, $2, 2, now())", gid, uid);
}
func queryGetSuggestion(db *sql.DB, g *suggests, uid int)  {
	rows, err := db.Query("select groups.groupid from groups left join usergroup on usergroup.groupid=groups.groupid where userid!=$1 and grouptype=1", uid);
	if err != nil {
		panicErr(err);
	}
	for rows.Next() {
		var t int;
		rows.Scan(&t);
		gg := new(group);
		queryGetGroup(db, gg, t);
		g.Suggest = append(g.Suggest, *gg);
	}
}
func queryCreateFriendList(db *sql.DB, uid int)  {
	db.Query("insert into groups (name, description, created, grouptype, updated) values ($1, $2, now(), 2, now())", "Friendlist", uid);
	var groupid int;
	db.QueryRow("select groupid from groups where grouptype=2 and description=$1", uid).Scan(&groupid);
	db.Query("insert into usergroup (groupid, userid, level, created) values($1, $2, 3, now())", groupid, uid);
}
func queryHasFriendList(db *sql.DB, uid int) bool  {
	_, err := db.Query("select groupid form usergroup where level=3 and userid=$1", uid)
	return err == nil
}
func queryAddFriend(db *sql.DB, uid1 int, uid2 int)  {
	var dd sql.NullString;
	err :=db.QueryRow("select groupid from usergroup where level=3 and userid=$1", uid2).Scan(&dd);
	panicErr(err);
	var ss sql.NullString;
	er := db.QueryRow("select groupid from usergroup where level=4 and groupid=$2 and userid=$1;",  uid1, dd.String).Scan(&ss);
//	panicErr(er);
	if er == nil{
		return
	} else{
		db.Query("insert into usergroup (groupid, userid, level, created) values($1, $2, 4, now())", dd.String, uid1);
	}
}
func queryRemoveFriend(db *sql.DB, uid1 int, uid2 int)  {
	var dd sql.NullString;
	err :=db.QueryRow("select groupid from usergroup where level=3 and userid=$1", uid2).Scan(&dd);
	panicErr(err);
	db.Query("delete from usergroup where groupid=$1 and userid=$2 and level=4", dd.String, uid1);
}
func queryFriendList(db *sql.DB, uid1 int, g *friendListPage)  {
	var dd sql.NullString;
	db.QueryRow("select groupid from usergroup where level=3 and userid=$1", uid1).Scan(&dd);
	rows, err := db.Query("select userid from usergroup where groupid=$1 and level=4", dd);
	panicErr(err);
	for rows.Next() {
		var t string;
		rows.Scan(&t);
		gg := new(user);
		queryGetUser(db, gg, t);
		g.UsersInfo = append(g.UsersInfo, *gg);
	}
}
func queryFriendListSugg(db *sql.DB, uid1 int, g *friendListPage)  {
	rows, err := db.Query("select userid from users where userid!=$1", uid1);
	panicErr(err);
	for rows.Next() {
		var t string;
		rows.Scan(&t);
		gg := new(user);
		queryGetUser(db, gg, t);
		g.UsersInfo = append(g.UsersInfo, *gg);
	}
}
func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}