package forum

// this file contains input structures

// data for new thread
type postMessage struct {
	Board   string
	boardid uint32 // cached
	UserID  *uint32
	PName   string
	Trip    string
	Email   string
	Title   string
	Message string
	Files   []fileContent
}

// data for new post in existing thread
type postData struct {
	postMessage
	ThreadID uint32
}

// when making new board
type boardData struct {
	Board          string // name
	Topic          string
	Description    string
	PageLimit      uint32
	ThreadsPerPage uint32
	AllowNewThread bool
	AllowFiles     bool
}
