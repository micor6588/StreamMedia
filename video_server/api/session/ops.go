package session

import (
	"StreamMedia/video_server/api/defs"
	"StreamMedia/video_server/api/utils"
	"golang-streaming/video_server/api/dbops"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func noInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// LoadSessionsFromDB 从数据库中加载Session
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

// GenerateNewSessionId 生生新的SessionId
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := noInMilli()
	ttl := ct + 30*60*1000 // Server side session valid time: 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

// IsSessionExpired 判断session是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
