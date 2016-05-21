
package news

import "time"

// external user info exported to forum/news/poll subsystems
// include fields only useful for everyone, skip ones which aren't going to be used elsewhere

type UserInfo struct {
	Text         string
	Score        float32
	CreationDate time.Time
	VisitCount   int
	Comment      string  
	Category     string
	LastModificationDate time.Time
	LastModificationAdmin string
	ThreadId     int
	}
