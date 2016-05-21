package forum

import (
	"../users"
	"time"
)

type boardInfo struct {
	Board       string
	Topic       string
	Description string
}

type frontPage struct {
	Boards []boardInfo
	// maybe add some stats or sth
}

type fileContent struct {
	Name     string // physical filename stored in server
	Original string // original filename user uploaded with
	Thumb    string // physical filename of thumbnail
}

func (f *fileContent) Valid() bool {
	return len(f.Name) > 0 && f.Name[0] != '/'
}

func (f *fileContent) VName() string {
	if f.Valid() {
		return f.Name
	} else {
		if len(f.Name) > 0 {
			return "[" + f.Name[1:] + "]"
		} else {
			return ""
		}
	}
}

type userIdent struct {
	Name  string          // name
	Trip  string          // tripcode
	Email string          // optional email field
	User  *users.UserInfo // user who posted message or nil. overrides other fields
}

func (i *userIdent) HasName() bool {
	return len(i.Name) > 0
}

func (i *userIdent) HasTrip() bool {
	return len(i.Trip) > 0
}

func (i *userIdent) HasEmail() bool {
	return len(i.Email) > 0
}

type postContent struct {
	PostID     uint32        // ID of post
	Title      string        // title of message
	UserIdent  userIdent     // identity of poster
	Date       time.Time     // exact time message was posted
	Message    string        // text of message
	Files      []fileContent // post files
	FMessage   string        // formatted message
	References []uint32      // references to post
}

type threadContent struct {
	ID      uint32         // thread ID
	OP      postContent    // OP
	Replies []postContent  // replies
	refmap  map[uint32]int // post ID -> reply id mapping
}

type threadInfo struct {
	ID       uint32
	Title    string
	OP       userIdent
	Replies  uint32
	Last     userIdent
	LastID   uint32
	LastDate time.Time
	Bump     time.Time
}

type boardPage struct {
	boardInfo
	// add extra info to board page
	Threads     []threadInfo
	Pages       []bool // used or not
	CurrentPage uint32
	Mod         bool // whether viewing in moderator mode or not
}

type threadPage struct {
	boardInfo
	Thread threadContent
	Mod    bool
}
