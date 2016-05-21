package users

import (
//"fmt"
	"../render"
	"net/http"
	"strconv"
	str "strings"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, pathi int) {
	// yet to be implemented
	http.Error(w, "501 not implemented", 501)
}
