package news

import (
	"../users"
	"time"
)

type articlesContent struct{
	NewsID uint32
	Title string
	UserIdent *users.UserInfo
	Date time.Time
	Text string
	Files []fileContent
}

type articlesList struct {
	Author       string
	Category       string
	Description string
	FullArticleLink string
}

type ArticlesFrontPage struct {
	Boards [] articlesList
	// maybe add some stats or sth
}

type fileContent struct {
	Name     string // physical filename stored in server
	Original string // original filename user uploaded with
}

/* type postContent struct {
	PostID     uint32 // ID of post
	Title      string // title of message
	UserIdent  userIdent
	Date       time.Time     // exact time message was posted
	Message    string        // text of message
	Files      []fileContent // post files
	FMessage   string        // formatted message
	References []uint32      // references to post
}
*/