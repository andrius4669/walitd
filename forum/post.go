package forum

import (
	"../dbacc"
	"../render"
	ss "../sessions"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	sc "strconv"
	"sync"
	"time"
)

const (
	maxImageSize = 8 << 20
	maxMusicSize = 50 << 20 // :^)
)

// nice place to also include file sizes
var allowedTypes = map[string]int64{
	"image/gif":  maxImageSize,
	"image/jpeg": maxImageSize,
	"image/png":  maxImageSize,
	"image/bmp":  maxImageSize,
	"audio/mpeg": maxMusicSize,
	"audio/ogg":  maxMusicSize,
	"audio/flac": maxMusicSize,
}

// add our own mime stuff since golang's parser erroreusly overwrites image/bmp with image/x-ms-bmp
func initMime() {
	mime.AddExtensionType(".bmp", "image/bmp")
	mime.AddExtensionType(".ogg", "audio/ogg")
	mime.AddExtensionType(".flac", "audio/flac")
}

// timestamps returned by this are guaranteed to be unique
var lastTimeMutex sync.Mutex
var lastTime int64 = 0

func uniqueTimestamp() int64 {
	lastTimeMutex.Lock()
	defer lastTimeMutex.Unlock()

	t := time.Now().UTC()
	unixnow := (t.Unix() * 1000) + ((t.UnixNano() / 1000000) % 1000)
	if unixnow > lastTime {
		lastTime = unixnow
		return unixnow
	} else {
		lastTime++
		return lastTime
	}
}

// new board creation
func handleNewBoard(w http.ResponseWriter, r *http.Request) {
	s := ss.GetUserSession(w, r)
	if s == nil {
		http.Error(w, "401 unauthorized: not logged in", 401)
		return
	}
	usi := new(ss.UserSessionInfo)
	ss.FillUserInfo(s, usi)
	if usi.Role < 2 {
		http.Error(w, "401 unauthorized: privilege too low", 401)
		return
	}

	r.ParseForm()
	d := new(boardData)
	bn, _ := r.Form["board"]
	if len(bn) > 0 {
		d.Board = bn[0]
	}
	tpc, _ := r.Form["topic"]
	if len(bn) > 0 {
		d.Topic = tpc[0]
	}
	dsc, _ := r.Form["desc"]
	if len(bn) > 0 {
		d.Description = dsc[0]
	}
	pl, _ := r.Form["pagelimit"]
	if len(pl) > 0 && pl[0] != "" {
		i, err := sc.Atoi(pl[0])
		if err == nil && i > 0 {
			d.PageLimit = uint32(i)
		}
	}
	tpp, _ := r.Form["threadsperpage"]
	if len(tpp) > 0 && tpp[0] != "" {
		i, err := sc.Atoi(tpp[0])
		if err == nil && i > 0 {
			d.ThreadsPerPage = uint32(i)
		}
	}
	yes, _ := r.Form["allownewthread"]
	if len(yes) > 0 && yes[0] == "yes" {
		d.AllowNewThread = true
	}
	yes, _ = r.Form["allowfiles"]
	if len(yes) > 0 && yes[0] == "yes" {
		d.AllowFiles = true
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	if !validateInputBoard(db, d) {
		fmt.Fprintf(w, "bad data (board failed to validate)")
		return
	}
	if sqlStoreBoard(db, d) {
		render.Execute(w, "boardmade", d)
	} else {
		fmt.Fprintf(w, "failed to make board")
	}
}

func newPostHandler(w http.ResponseWriter, r *http.Request, threadid uint32, files []fileContent, board string) {
	p := new(postData)
	p.ThreadID = threadid
	p.Board = board
	nn, _ := r.Form["name"]
	if len(nn) > 0 {
		p.PName = nn[0]
	}
	ee, _ := r.Form["email"]
	if len(ee) > 0 {
		p.Email = ee[0]
	}
	if p.Email == "sage" {
		p.bump = false
	} else {
		p.bump = true
	}
	tt, _ := r.Form["title"]
	if len(tt) > 0 {
		p.Title = tt[0]
	}
	mm, _ := r.Form["message"]
	if len(mm) > 0 {
		p.Message = mm[0]
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	for i := range files {
		p.Files = append(p.Files, files[i])
	}

	if !validateInputPost(db, p) {
		fmt.Fprintf(w, "bad data (post failed to validate)")
		return
	}
	if sqlStorePost(db, p) {
		render.Execute(w, "postmade", p)
	} else {
		fmt.Fprintf(w, "failed to make post")
	}
}

func newThreadHandler(w http.ResponseWriter, r *http.Request, files []fileContent, board string) {
	p := new(postMessage)
	p.Board = board
	nn, _ := r.Form["name"]
	if len(nn) > 0 {
		p.PName = nn[0]
	}
	ee, _ := r.Form["email"]
	if len(ee) > 0 {
		p.Email = ee[0]
	}
	tt, _ := r.Form["title"]
	if len(tt) > 0 {
		p.Title = tt[0]
	}
	mm, _ := r.Form["message"]
	if len(mm) > 0 {
		p.Message = mm[0]
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	for i := range files {
		p.Files = append(p.Files, files[i])
	}

	if !validateInputThread(db, p) {
		fmt.Fprintf(w, "bad data (thread failed to validate)")
		return
	}

	if sqlStoreThread(db, p) {
		render.Execute(w, "threadmade", p)
	} else {
		fmt.Fprintf(w, "failed to make post")
	}
}

// new thread or post creation
func handlePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		fmt.Fprintf(w, "bad request")
		return
	}

	var board string
	bn, _ := r.Form["board"]
	if len(bn) > 0 {
		board = bn[0]
	}

	var tid uint32
	thr, _ := r.Form["thread"]
	if len(thr) > 0 {
		i, _ := sc.Atoi(thr[0])
		tid = uint32(i)
	}

	var files []fileContent
	if r.MultipartForm.File != nil {
		fhs := r.MultipartForm.File["file"]
		for i := range fhs {
			f, err := fhs[i].Open()
			if err != nil {
				continue
			}
			defer f.Close()

			size, err := f.Seek(0, os.SEEK_END)
			if err != nil {
				http.Error(w, fmt.Sprintf("500 internal server error: %s", err), 500)
				return
			}
			_, err = f.Seek(0, os.SEEK_SET)
			if err != nil {
				http.Error(w, fmt.Sprintf("500 internal server error: %s", err), 500)
				return
			}

			var p fileContent

			ext := filepath.Ext(fhs[i].Filename)
			mt := mime.TypeByExtension(ext)
			if mt != "" {
				mt, _, _ = mime.ParseMediaType(mt)
			}
			maxSize, ok := allowedTypes[mt]
			if !ok {
				http.Error(w, "file type not allowed", 403) // 403 Forbidden
				return
			}

			if size > maxSize {
				http.Error(w, "file too big", 403) // 403 Forbidden
				return
			}

			fname := sc.FormatInt(uniqueTimestamp(), 10) + ext
			os.MkdirAll(serveSrcPathDir(board), 777)
			fullname := serveSrcPath(board, fname)
			tmpname := serveSrcPath(board, ".tmp."+fname)
			nf, err := os.OpenFile(tmpname, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				http.Error(w, fmt.Sprintf("500 internal server error: %s", err), 500)
				return
			}
			io.Copy(nf, f)
			nf.Close()
			os.Rename(tmpname, fullname) // atomic :^)

			p.Name = fname
			p.Original = fhs[i].Filename

			tname, err := makeThumb(fullname, fname, board, ext, mt, tid == 0)
			if err != nil {
				fmt.Printf("error generating thumb for %s: %s\n", fname, err)
			}
			p.Thumb = tname

			files = append(files, p)
		}
	}

	if tid != 0 {
		newPostHandler(w, r, tid, files, board)
	} else {
		newThreadHandler(w, r, files, board)
	}
}

func handleDelBoard(w http.ResponseWriter, r *http.Request) {
	s := ss.GetUserSession(w, r)
	if s == nil {
		http.Error(w, "401 unauthorized: not logged in", 401)
		return
	}
	usi := new(ss.UserSessionInfo)
	ss.FillUserInfo(s, usi)
	if usi.Role < 2 {
		http.Error(w, "401 unauthorized: privilege too low", 401)
		return
	}

	r.ParseForm()
	var board string
	bn, _ := r.Form["board"]
	if len(bn) > 0 {
		board = bn[0]
	}
	if board == "" {
		fmt.Fprintf(w, "bad data (board name is empty)")
		return
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	if sqlDeleteBoard(db, board) {
		fmt.Fprintf(w, "board deleted")
	} else {
		fmt.Fprintf(w, "failed to delete board")
	}
}

func handleDelPost(w http.ResponseWriter, r *http.Request) {
	s := ss.GetUserSession(w, r)
	if s == nil {
		http.Error(w, "401 unauthorized: not logged in", 401)
		return
	}
	usi := new(ss.UserSessionInfo)
	ss.FillUserInfo(s, usi)
	if usi.Role < 2 {
		http.Error(w, "401 unauthorized: privilege too low", 401)
		return
	}

	r.ParseForm()

	var board string
	var post uint32

	bn, _ := r.Form["board"]
	if len(bn) > 0 {
		board = bn[0]
	}
	pst, _ := r.Form["post"]
	if len(pst) > 0 {
		i, _ := sc.Atoi(pst[0])
		post = uint32(i)
	}
	if board == "" {
		fmt.Fprintf(w, "bad data (board name is empty)")
		return
	}
	if post == 0 {
		fmt.Fprintf(w, "bad data (post undefined)")
		return
	}

	db := dbacc.OpenSQL()
	defer db.Close()

	if sqlDeletePost(db, board, post) {
		fmt.Fprintf(w, "post deleted")
	} else {
		fmt.Fprintf(w, "failed to delete post")
	}
}
