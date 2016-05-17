
package news

import "time"

// external user info exported to forum/news/poll subsystems
// include fields only useful for everyone, skip ones which aren't going to be used elsewhere

type UserInfo struct {
	Text         string
	Score        Float32    
	CreationDate time.time    
	VisitCount   Int 
	Comment      string  
	Category     string
	LastModificationDate time.time
	LastModificationAdmin string
	ThreadId     int
	}
