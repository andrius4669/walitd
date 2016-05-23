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
	rows, err := db.Query("SELECT article_id, article_text, score, visit_count, description, category, author, upload_date, article_name FROM article ORDER BY upload_date DESC")
	panicErr(err)

	for rows.Next() {
		var b articlesList
		rows.Scan(&b.ID, &b.Article, &b.Score, &b.Visit_Count, &b.Description, &b.Category, &b.Author, &b.UploadDate, &b.Name)

		//fmt.Printf("%v \n", rows);
		err := db.QueryRow("SELECT username FROM users WHERE userid=$1", b.Author).Scan(&b.AuthorName)
		//fmt.Printf("%v \n", b.Name);
		panicErr(err)
		//fmt.Printf("%v \n", b.AuthorName);

		p.Boards = append(p.Boards, b)
	}
}

func queryArticle(db *sql.DB, p *articlesList, id int) {
	err := db.QueryRow("SELECT article_text, score, visit_count, description, category, author, upload_date FROM article WHERE article_id=$1", id).Scan(&p.Article, &p.Score, &p.Visit_Count, &p.Description, &p.Category, &p.Author, &p.UploadDate)
	panicErr(err)
	updateVisits(db, id)
}

func updateVisits(db *sql.DB, id int) {
	db.QueryRow("UPDATE article SET visit_count = visit_count + 1 WHERE article_id=$1", id)
	//panicErr(err)
}
func createArticle(db *sql.DB, p *articlesList) {
	//db.QueryRow("INSERT INTO article (article, score, upload_date, visit_count, description, category, last_modification_date, last_modification_admin, thread_id, author, articleName) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", p.Article, p.Score, p.UploadDate, 0, p.Description, p.Category, p.UploadDate, p.Author, 0, p.Author, p.Name)
	_, err := db.Query("INSERT INTO article (article_text, score, upload_date, visit_count, description, category, last_modification_date, last_modification_admin, thread_id, author, article_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", p.Article, p.Score, p.UploadDate, 0, p.Description, p.Category, p.UploadDate, p.Author, 1, p.Author, p.Name)
	panicErr(err);
	//panicErr(err)INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country)
	//VALUES ('Cardinal','Tom B. Erichsen','Skagen 21','Stavanger','4006','Norway');
}