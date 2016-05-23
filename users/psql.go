package users
import (
"fmt"
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
	rows, err := db.Query("Select sender, reciever, message, created from messages where sender=$1;", id);
	panicErr(err);
	for rows.Next() {
		var t message
		rows.Scan(&t.Sender, &t.Reciever, &t.Text, &t.Created);
		m.Sent = append(m.Sent, t);
	}
	rows, err = db.Query("Select sender, reciever, message, created from messages where reciever=$1;", id);
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
	err := db.QueryRow("Select name, description, created, updated from groups where groupid=$1 and grouptype=1;", id).Scan(&g.Name, &desc, &g.Created, &g.Updated);
	if (desc.Valid){
		g.Description = desc.String;
	}
	if err != nil {
		panicErr(err)
		return err;
	}
	pp := db.QueryRow("select users.userid, users.username from usergroup left join users on users.userid=usergroup.userid where level=1 and groupid=$1", id).Scan(&g.Owner, &g.OwnerName);
	panicErr(err)
	return pp;
}
func queryGetGroupByName(db *sql.DB, g *group, id string){
	err := db.QueryRow("Select name, description, created, updated from groups where name=$1 and grouptype=1;", id).Scan(&g.Name, &g.Description, &g.Created, &g.Updated);
	panicErr(err);
}
func queryLogin(db *sql.DB, ul *loginInfo, u *user){
	var id string;
	err := db.QueryRow("select id from users where username=$1 and password=$2", ul.Username, ul.Pass).Scan(&id);
	panicErr(err);
	if id != "" {
		queryGetUser(db, u, id);
	}
}
func queryCreateGroup(db *sql.DB, g *group)  {
	_, err := db.Query("insert into groups (groupid, name, description, created, grouptype, updated) values (0, $1, $2, now(), 1, now())", g.Name, g.Description);
	panicErr(err);
}
func queryAddToGroup(db *sql.DB, groupid int, level int, userid int){
	_, err := db.Query("insert into usergroups (gruopid, userid, level, created) values($3, $1, $2, now())", userid, level, groupid);
	panicErr(err);
}
func queryAddMessage(db *sql.DB, m *messageForm)  {
	_, err := db.Query("insert into messages (id, sender, reciever, mesage, created) values(0, $1, $2, $3, now())", m.Sender, m.To, m.Message);
	panicErr(err);
}
func queryAddUser(db *sql.DB, u *userForm){
	_, err := db.Query("insert into users (username, password, firstname, lastname, role, gender, created, updated, photocreated, active, email) values($1, $2, $3, $4, $5, $6, now(), now(), now(), 1, $7);", u.Username, u.Pass, u.FirstName, u.SecondName, 1, u.Gender, u.Email);
	panicErr(err);
}
func queryUpdateUser(db *sql.DB, u *user){
	fmt.Printf("%v \n", u);
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
func queryGetFriendList(db *sql.DB, f *friendListPage, userid int){
	var ss string;
	err := db.QueryRow("Select usergroups.groupid from usergroups left join groups on usergroups.groupid=groups.groupid where usergroups.userid=$1 and groups.grouptype=2 and usergroups.level=1", userid).Scan(&ss);
	panicErr(err);
	rows, err2 := db.Query("Select userid from usergroups where groupid=$1 and userid!=$2", ss, userid);
	panicErr(err2);
	for rows.Next() {
		var t string
		rows.Scan(&t);
		u := new(user);
		queryGetUser(db, u, t);
		f.UsersInfo = append(f.UsersInfo, *u);
	}
}
func queryGetGroupList(db *sql.DB,g *groupsPage, userid int) {
	rows, err := db.Query("Select groupid from usergroups left join groups on usergroups.groupid=groups.groupid where usergroups.userid=$1 and groups.grouptype=1", userid);
	panicErr(err);
	for rows.Next() {
		var t int;
		rows.Scan(&t);
		gg := new(group);
		queryGetGroup(db, gg, t);

		g.GroupsInfo = append(g.GroupsInfo, *gg);
	}
}
func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}