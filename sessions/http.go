package sessions

import (
	"net/http"
	"time"
)

type  UserSessionInfo struct {
	Uid  uint32
	Role uint32
}

// when making new session, use this
func MakeUserSession(w http.ResponseWriter, r *http.Request, si *UserSessionInfo) *SessionStore {
	cookie, _ := r.Cookie(manager.cookieName)
	var sid string
	if cookie != nil && cookie.Value != "" {
		sid = cookie.Value
	} else {
		sid = manager.newSessionID()
	}
	s := manager.SessionUpdateOrNew(sid)
	s.val["uid"] = si.Uid
	s.val["role"] = si.Role
	exp := time.Now().Add(time.Duration(manager.maxlifetime) * time.Second)
	sc := http.Cookie{Name: manager.cookieName, Value: sid, Expires: exp}
	http.SetCookie(w, &sc)
	return s
}

// when checking if active session exists, use this
func GetUserSession(w http.ResponseWriter, r *http.Request) *SessionStore {
	cookie, _ := r.Cookie(manager.cookieName)
	if cookie == nil || cookie.Value == "" {
		return nil
	}
	sid := cookie.Value
	s := manager.SessionUpdate(sid)
	if s == nil {
		return nil
	}
	exp := time.Now().Add(time.Duration(manager.maxlifetime) * time.Second)
	sc := http.Cookie{Name: manager.cookieName, Value: sid, Expires: exp}
	http.SetCookie(w, &sc)
	return s
}

// extracts userinfo from session
func FillUserInfo(s *SessionStore, usi *UserSessionInfo) {
	usi.Uid  = s.val["uid"].(uint32)
	usi.Role = s.val["uid"].(uint32)
}
