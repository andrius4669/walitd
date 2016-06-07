package news

import (
	"../users"
	"time"
)

type articlesContent struct{
	NewsID uint32
	Title string
	UserIdent users.UserInfo
	Date time.Time
	Text string
	Files []fileContent
}

type articlesList struct {
	ID int
	Article string
	Name string
	Score int
	Visit_Count int
	Description string
	Category string
	Author       int
	AuthorName string
	UploadDate string
	FullArticleLink string
	Tags string
}

type ArticlesFrontPage struct {
	Boards [] articlesList
	// maybe add some stats or sth
}

type voteInfo struct {
	user_id int
	article_id int
	vote string
}

type fileContent struct {
	Name     string // physical filename stored in server
	Original string // original filename user uploaded with
}
