package users
import (
//"fmt"
//	"encoding/json"
	"database/sql"
//	"strconv"
)

func queryGetUser(db *sql.DB, u *user, id string){
	err := db.QueryRow("SELECT userid, username, email, name, surname, role, birthday, city, country, telephone, gender, description, created, updated, photo, photocreated FROM users where userid=$1", id).Scan(&u.Userid, &u.Username, &u.Email, &u.FirstName, &u.SecondName, &u.Role, &u.Birthday, &u.City, &u.Country, &u.Telephone, &u.Gender, &u.Description, &u.Created, &u.Updated, &u.Picture, &u.PictureCreated);
	panicErr(err);
}
func queryGetUserByUsername(db *sql.DB,u *user, username string){
	err := db.QueryRow("SELECT userid, username, email, name, surname, role, birthday, city, country, telephone, gender, description, created, updated, photo, photocreated FROM users where userid=$1", username).Scan(&u.Userid, &u.Username, &u.Email, &u.FirstName, &u.SecondName, &u.Role, &u.Birthday, &u.City, &u.Country, &u.Telephone, &u.Gender, &u.Description, &u.Created, &u.Updated, &u.Picture, &u.PictureCreated);
	panicErr(err);
}
func queryGetMessages(db *sql.DB, m *messages, id int){
	rows, err := db.Query("Select sender, reciever, message, created from messages where sender=$1", id);
	panicErr(err);
	for rows.Next() {
		var t message
		rows.Scan(&t.Sender, &t.Reciever, &t.Text, &t.Created);
		m.Sent = append(m.Sent, t);
	}
	rows, err = db.Query("Select sender, reciever, message, created from messages where reciever=$1", id);
	panicErr(err);
	for rows.Next() {
		var t message
		rows.Scan(&t.Sender, &t.Reciever, &t.Text, &t.Created);
		m.Recieved = append(m.Recieved, t);
	}
}
func queryGetGroup(db *sql.DB, g *group, id int){
	err := db.QueryRow("Select name, description, created, updated from groups where id=$1 and grouptype=1", id).Scan(&g.Name, &g.Description, &g.Created, &g.Updated);
	panicErr(err);
}
func queryGetGroupByName(db *sql.DB, g *group, id string){
	err := db.QueryRow("Select name, description, created, updated from groups where name=$1 and grouptype=1", id).Scan(&g.Name, &g.Description, &g.Created, &g.Updated);
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
//func queryAddToGroup(db *sql.DB, groupid int, level int){
//	err := db.QueryRow("insert into usergroups ()")
//	panicErr(err);
//}
func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}