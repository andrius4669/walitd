package files

import (
	"net/http"
	"os"
	"time"
)

func ServeFile(w http.ResponseWriter, r *http.Request, fname string) bool {
	f, err := os.Open(fname)
	if err != nil {
		return false
	}
	defer f.Close()

	fi, err := f.Stat()
	if err == nil {
		http.ServeContent(w, r, fname, fi.ModTime(), f)
	} else {
		http.ServeContent(w, r, fname, time.Time{}, f)
	}
	return true
}

func ServeFileOr404(w http.ResponseWriter, r *http.Request, fname string) bool {
	ok := ServeFile(w, r, fname)
	if !ok {
		http.NotFound(w, r)
	}
	return ok
}
