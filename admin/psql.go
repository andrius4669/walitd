package admin

import (
	"database/sql"
	//"fmt"
)

func queryDeleteArticle(db *sql.DB, name string) {
	var id int
	err := db.QueryRow("SELECT article_id FROM article WHERE article_name=$1;", name).Scan(&id);
	panicErr(err);


	db.QueryRow("DELETE FROM article WHERE article_name=$1", name)

	db.QueryRow("DELETE FROM article_tags WHERE article_id=$1", id)
	db.QueryRow("DELETE FROM score WHERE article_id=$1", id)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getArticle(db *sql.DB, p *articlesList, name string){
	err := db.QueryRow("SELECT article_text, score, visit_count, description, category, author, upload_date, article_name FROM article WHERE article_name=$1", name).Scan(&p.Article, &p.Score, &p.Visit_Count, &p.Description, &p.Category, &p.Author, &p.UploadDate, &p.Name)
	panicErr(err)
}
func updateArticle(db *sql.DB, p *articlesList){
	db.QueryRow("UPDATE article SET article_text=$2, description=$3, article_name=$4 WHERE article_id=$1", p.ID, p.Article, p.Description, p.Name)
	//panicErr(err)
}

func getArticleID(db *sql.DB, p *articlesList, id *int) int {
	err := db.QueryRow("SELECT article_id FROM article WHERE article_name=$1;", p.Name).Scan(&id);
	panicErr(err);
	return *id
}
