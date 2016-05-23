package sessions

import (
	"net/http"
	"net/url"
	"time"
)

type UserSessionInfo struct {
	Uid  uint32
	Role uint32
}

// when making new session, use this
func MakeUserSession(w http.ResponseWriter, r *http.Request, si *UserSessionInfo) *SessionStore {
	cookie, _ := r.Cookie(manager.cookieName)
	var sid string
	if cookie != nil && cookie.Value != "" {
		sid, _ = url.QueryUnescape(cookie.Value)
	}
	if sid == "" {
		sid = manager.newSessionID()
	}
	s := manager.SessionUpdateOrNew(sid)
	s.val["uid"] = si.Uid
	s.val["role"] = si.Role
	exp := time.Now().Add(time.Duration(manager.maxlifetime) * time.Second)
	sc := http.Cookie{
		Name:     manager.cookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		Expires:  exp,
		MaxAge:   int(manager.maxlifetime),
	}
	http.SetCookie(w, &sc)
	return s
}

// when checking if active session exists, use this
func GetUserSession(w http.ResponseWriter, r *http.Request) *SessionStore {
	cookie, _ := r.Cookie(manager.cookieName)
	if cookie == nil || cookie.Value == "" {
		return nil
	}
	sid, _ := url.QueryUnescape(cookie.Value)
	if sid == "" {
		return nil
	}
	s := manager.SessionUpdate(sid)
	if s == nil {
		return nil
	}
	exp := time.Now().Add(time.Duration(manager.maxlifetime) * time.Second)
	sc := http.Cookie{
		Name:     manager.cookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		Expires:  exp,
		MaxAge:   int(manager.maxlifetime),
	}
	http.SetCookie(w, &sc)
	return s
}

// for logging out
func pruneUserSession(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(manager.cookieName)
	if cookie == nil || cookie.Value == "" {
		return // no session
	}
	sid, _ := url.QueryUnescape(cookie.Value)
	if sid != "" {
		manager.SessionPrune(sid)
	}
	exp := time.Now()
	sc := http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  exp,
		MaxAge:   -1,
	}
	http.SetCookie(w, &sc)
}

// extracts userinfo from session
func FillUserInfo(s *SessionStore, usi *UserSessionInfo) {
	usi.Uid = s.val["uid"].(uint32)
	usi.Role = s.val["uid"].(uint32)
}
