package forum

/*
 * most of stuff in this file isn't really about serving itself,
 * but about validation & providing right paths to files.ServeFileOr404 function
 */

import (
	cfg "../configmgr"
	"../files"
	"net/http"
	"path"
)

// board does not contain /, it's filtered out by caller
// file can not contain / too
func serveSrcPathDir(board string) string {
	// board cannot be . or .., file cannot start with . (we will hide tmp files that way)
	if board == "." || board == ".." {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	sdir, _ := cfg.GetOption("forum.srcdir")
	return fdir + "/" + board + "/" + sdir
}

// board does not contain /, it's filtered out by caller
// file can not contain / too
func serveSrcPath(board, file string) string {
	// board cannot be . or .., file cannot start with . (we will hide tmp files that way)
	if board == "." || board == ".." {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	sdir, _ := cfg.GetOption("forum.srcdir")
	return fdir + "/" + board + "/" + sdir + "/" + file
}

func serveSrc(w http.ResponseWriter, r *http.Request, board, file string) {
	file = serveSrcPath(board, file)
	if file == "" {
		http.NotFound(w, r)
		return
	}
	files.ServeFileOr404(w, r, file)
}

func serverThumbPathDir(board string) string {
	if board == "." || board == ".." {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	tdir, _ := cfg.GetOption("forum.thumbdir")
	return fdir + "/" + board + "/" + tdir
}

// board does not contain /, it's filtered out by caller
// file could contain /
func serverThumbPath(board, file string) string {
	if board == "." || board == ".." || file[0] == '.' {
		return ""
	}

	cfile := path.Clean("/" + file)
	if cfile == "" || cfile == "/" {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	tdir, _ := cfg.GetOption("forum.thumbdir")
	return fdir + "/" + board + "/" + tdir + cfile
}

func serveThumb(w http.ResponseWriter, r *http.Request, board, file string) {
	file = serverThumbPath(board, file)
	if file == "" {
		http.NotFound(w, r)
		return
	}
	files.ServeFileOr404(w, r, file)
}

// file can contain / and it's OK
func serveStaticPath(file string) string {
	if file[0] == '.' {
		return ""
	}
	cfile := path.Clean("/" + file)
	if cfile == "" || cfile == "/" {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	sdir, _ := cfg.GetOption("forum.staticdir")
	return fdir + "/" + sdir + cfile
}

func serveStatic(w http.ResponseWriter, r *http.Request, file string) {
	file = serveStaticPath(file)
	if file == "" {
		http.NotFound(w, r)
		return
	}
	files.ServeFileOr404(w, r, file)
}

// file can contain / and it's OK
func serveBoardStaticPath(board, file string) string {
	if board == "." || board == ".." || file[0] == '.' {
		return ""
	}
	cfile := path.Clean("/" + file)
	if cfile == "" || cfile == "/" {
		return ""
	}

	fdir, _ := cfg.GetOption("forum.filedir")
	sdir, _ := cfg.GetOption("forum.boardstaticdir")
	return fdir + "/" + board + "/" + sdir + cfile
}

func serveBoardStatic(w http.ResponseWriter, r *http.Request, board, file string) {
	file = serveBoardStaticPath(board, file)
	if file == "" {
		http.NotFound(w, r)
		return
	}
	files.ServeFileOr404(w, r, file)
}
