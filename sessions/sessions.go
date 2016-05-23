package sessions

import (
	"time"
	"sync"
	"container/list"
)

// stores sessions

type SessionStore struct {
	sid string    // unique session ID
	st  time.Time // access time
	val map[string]interface{}
}

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	maxlifetime int64
	sessions    map[string]*list.Element
	list        *list.List
}

// checks gc
func (m *Manager) checkGC() {
	for {
		el := m.list.Back()
		if el == nil {
			break
		}
		if el.Value.(*SessionStore).st.Unix() > (time.Now().Unix() + m.maxlifetime) {
			m.list.Remove(el)
			delete(m.sessions, el.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// creates new
func (m *Manager) SessionInit(sid string) *SessionStore {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.checkGC()

	v := make(map[string]interface{}, 0)
	ns := &SessionStore{sid: sid, st: time.Now(), val: v}
	el := m.list.PushFront(ns)
	m.sessions[sid] = el

	return ns
}

// gets (or returns nil if not exists)
func (m *Manager) SessionGet(sid string) *SessionStore {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.checkGC()

	s, ok := m.sessions[sid] // gets... or not
	if ok {
		return s.Value.(*SessionStore)
	}
	return nil
}

// gets and bumps in list if exists
func (m *Manager) SessionUpdate(sid string) *SessionStore {
	m.lock.Lock()
	defer m.lock.Unlock()

	s, ok := m.sessions[sid]
	if ok {
		s.Value.(*SessionStore).st = time.Now() // update time
		m.list.MoveToFront(s)

		m.checkGC()

		return s.Value.(*SessionStore)
	}

	m.checkGC()

	return nil
}

var manager = &Manager{}

func init() {
	manager.cookieName = "session"
	manager.maxlifetime = 3600
	manager.sessions = make(map[string]*list.Element)
	manager.list = list.New()
}

func GetManager() *Manager {
	return manager
}
