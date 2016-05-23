package news

import (
	//"fmt"
	//	"encoding/json"
	//	"strconv"
	//"encoding/json"
	"database/sql"
)

//func queryGetArticleList(db *sql.DB, u *user, id string){
//	err := db.QueryRow("SELECT * FROM article where article_id=$1;", id).Scan(&u.Userid, &u.Username, &u.Email, &u.FirstName, &u.SecondName, &u.Role, &u.Birthday, &u.City, &u.Country, &u.Telephone, &u.Gender, &u.Description, &u.Created, &u.Updated, &u.Picture, &u.PictureCreated);
//	panicErr(err);
//}
//ArticleList
func queryArticlesList(db *sql.DB, p *ArticlesFrontPage) {
	rows, err := db.Query("SELECT article, score, visit_count, description, category, author, upload_date FROM article")
	panicErr(err)

	for rows.Next() {
		var b articlesList
		rows.Scan(&b.Name, &b.Score, &b.Visit_Count, &b.Description, &b.Category, &b.Author, &b.UploadDate)

		rows2, err := db.Query("SELECT username FROM users WHERE userid=$1",b.Name)
		panicErr(err)
		rows2.Scan(&b.AuthorName)

		p.Boards = append(p.Boards, b)
	}
}