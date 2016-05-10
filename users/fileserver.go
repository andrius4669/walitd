package users

import (
	cfg "../configmgr"
	"../files"
	"net/http"
	//"path"
)

// file shall not contain /, it should be filtered out by caller
func serveAvatarPath(file string) string {
	if file[0] == '.' {
		return ""
	}
	ad, _ := cfg.GetOption("users.avatardir")
	return ad + "/" + file
}

func serveAvatar(w http.ResponseWriter, r *http.Request, file string) {
	file = serveAvatarPath(file)
	if file == "" {
		http.NotFound(w, r)
		return
	}
	files.ServeFileOr404(w, r, file)
}
