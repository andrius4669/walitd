package forum

import (
	"fmt"
	"net/http"
)

// TODO(andrius)

func renderBoardList(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 board list not implemented", 501)
}

func renderBoardListModPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 board list mod page not implemented", 501)
}

func renderBoardModPage(w http.ResponseWriter, r *http.Request, board string) {
	http.Error(w, fmt.Sprintf("501 board %s mod page not implemented", board), 501)
}

func renderBoardPage(w http.ResponseWriter, r *http.Request, board string, page uint32, mod bool) {
	http.Error(w, fmt.Sprintf("501 board %s page %d (mod: %t) not implemented", board, page, mod), 501)
}

func renderThread(w http.ResponseWriter, r *http.Request, board string, thread string, mod bool) {
	http.Error(w, fmt.Sprintf("501 board %s thread %s (mod: %t) not implemented", board, thread, mod), 501)
}
