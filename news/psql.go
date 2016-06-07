package news

import (

	fm "../forum"
	"database/sql"
	"strings"
	"bytes"
	"fmt"
	//"math"
)


func queryArticlesList(db *sql.DB, p *ArticlesFrontPage) {
	rows, err := db.Query("SELECT article_id, article_text, score, visit_count, description, category, author, upload_date, article_name FROM article ORDER BY upload_date DESC")
	panicErr(err)

	for rows.Next() {
		var b articlesList
		rows.Scan(&b.ID, &b.Article, &b.Score, &b.Visit_Count, &b.Description, &b.Category, &b.Author, &b.UploadDate, &b.Name)
		getTags(db, b.ID, &b.Tags);
		err := db.QueryRow("SELECT username FROM users WHERE userid=$1", b.Author).Scan(&b.AuthorName)
		panicErr(err)

		p.Boards = append(p.Boards, b)
	}
}

func queryArticlesSearchList(db *sql.DB, p *ArticlesFrontPage, search []string) {

	var b articlesList
	for _,element := range search {
		element2 := strings.TrimSpace(element)
		rows, err := db.Query("SELECT article_id FROM article_tags WHERE tag_name=$1", element2)
		panicErr(err)
		for rows.Next() {
			var id int
			rows.Scan(&id)
			err := db.QueryRow("SELECT article_id, article_text, score, visit_count, description, category, author, upload_date, article_name FROM article WHERE article_id=$1 ORDER BY upload_date DESC", id).Scan(&b.ID, &b.Article, &b.Score, &b.Visit_Count, &b.Description, &b.Category, &b.Author, &b.UploadDate, &b.Name)
			panicErr(err)
			getTags(db, b.ID, &b.Tags);
			err2 := db.QueryRow("SELECT username FROM users WHERE userid=$1", b.Author).Scan(&b.AuthorName)
			panicErr(err2)

			p.Boards = append(p.Boards, b)
			}
		}
}



func queryArticle(db *sql.DB, p *articlesList, id int) {
	err := db.QueryRow("SELECT article_text, score, visit_count, description, category, author, upload_date FROM article WHERE article_id=$1", id).Scan(&p.Article, &p.Score, &p.Visit_Count, &p.Description, &p.Category, &p.Author, &p.UploadDate)
	getTags(db, id, &p.Tags);
	panicErr(err)
	updateVisits(db, id)
}

func updateVisits(db *sql.DB, id int) {
	db.QueryRow("UPDATE article SET visit_count = visit_count + 1 WHERE article_id=$1", id)
}
func createArticle(db *sql.DB, p *articlesList) {
	_, err := db.Query("INSERT INTO article (article_text, score, upload_date, visit_count, description, category, last_modification_date, last_modification_admin, thread_id, author, article_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", p.Article, p.Score, p.UploadDate, 0, p.Description, p.Category, p.UploadDate, p.Author, 1, p.Author, p.Name)
	panicErr(err);
	db.QueryRow("SELECT article_id FROM article WHERE article_name=$1;", p.Name).Scan(&p.ID);

	db.Query("INSERT INTO score (article_id, likes, dislikes) values($1, 0, 0)", p.ID)

	//var data boardData
	//data.Board = p.ID
	//data.Topic = p.ID
	//data.Desription = "Article description"
	//data.PageLimit = 100
	//data.ThreadsPerPage = 10
	//data.AllowNewThread = true
	//data.AllowFiles = true
	//fm.handleNewBoardArticle(data)
}
func addTags(db *sql.DB, tags []string, id int) {
	for _,element := range tags {
		element2 := strings.TrimSpace(element)
		addTag(db, element2, id)
	}
}
func addTag(db *sql.DB, tag string, id int) {
	_, err := db.Query("INSERT INTO article_tags (article_id, tag_name) VALUES($1, $2)", id, tag)
	panicErr(err);
}

func getArticleID(db *sql.DB, p *articlesList, tags[]string, id int){
	err := db.QueryRow("SELECT article_id FROM article WHERE article_name=$1;", p.Name).Scan(&id);
	addTags(db, tags, id)
	panicErr(err);
}

func getTags(db *sql.DB, id int, tags *string){

	rows, err := db.Query("SELECT tag_name FROM article_tags WHERE article_id=$1", id)
	panicErr(err)
	var buffer bytes.Buffer
	for rows.Next() {
		var tagsString string
		rows.Scan(&tagsString)
		buffer.WriteString(tagsString)
	}
	*tags = buffer.String()

}

func createTags(db *sql.DB, tags []string){
	for _,element := range tags {
		var numrepl int
		element2 := strings.TrimSpace(element)
		err := db.QueryRow("SELECT COUNT(*) FROM tags WHERE name=$1", element2).Scan(&numrepl)
		panicErr(err)
		if numrepl > 0 {
			db.QueryRow("UPDATE tags SET use_count = use_count + 1 WHERE name=$1", element2)
		} else {
			_, err := db.Query("INSERT INTO tags (name, use_count, author, creation_date) VALUES($1, $2, $3, now())", element2, 0, 0)
			panicErr(err);
		}

	}
}

func voteForArticle(db *sql.DB, articleID int , userID int, vote string){

	var b voteInfo
	var err error
	err = db.QueryRow("SELECT  article_id, user_id, vote FROM user_vote WHERE user_id=$1 AND article_id=$2;", userID, articleID).
	Scan(&b.article_id, &b.user_id, &b.vote);

	trimedOldVote := strings.TrimSpace(b.vote)

	if(err == sql.ErrNoRows){
		db.Query("INSERT INTO user_vote (article_id, user_id, vote) values($1, $2, $3)", articleID, userID, vote)
		if vote == "like" {
			db.QueryRow("UPDATE score SET like = 1 WHERE article_id=$1", articleID)
		} else if vote == "dislike" {
			db.QueryRow("UPDATE score SET dislike = 1 WHERE article_id=$1", articleID)
		}


	} else if(trimedOldVote != vote){

		if vote == "like" {
			fmt.Printf("VOTING FOR ARTIICLE - LIKE OLD \n");
			db.QueryRow("UPDATE score SET likes = likes + 1 WHERE article_id=$1", articleID)
			db.QueryRow("UPDATE user_vote SET vote = $1 WHERE article_id=$2 AND user_id=$3",vote , articleID, userID)
			db.QueryRow("UPDATE score SET dislikes = dislikes - 1 WHERE article_id=$1", articleID)

		} else if vote == "dislike" {
			fmt.Printf("VOTING FOR ARTIICLE - DISLIKE OLD \n");
			db.QueryRow("UPDATE score SET likes = likes - 1 WHERE article_id=$1", articleID)
			db.QueryRow("UPDATE user_vote SET vote = $1 WHERE article_id=$2 AND user_id=$3",vote , articleID, userID)
			db.QueryRow("UPDATE score SET dislikes = dislikes + 1 WHERE article_id=$1", articleID)

		}
	}
	countScore(db, articleID)

}
func countScore(db *sql.DB, articleID int){
	var like int
	var dislike int
	db.QueryRow("SELECT  likes, dislikes FROM score WHERE article_id=$1", articleID).
	Scan(&like, &dislike);
	var score float32
	flike := float32(like)
	fdislike := float32(dislike)
 	score = flike/(flike+fdislike) * 100
	db.QueryRow("UPDATE article SET score = $1 WHERE article_id=$2", score,articleID)

}










