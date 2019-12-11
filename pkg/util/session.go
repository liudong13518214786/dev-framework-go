package util

import (
	"dev-framework-go/pkg/cache"
)

type SessionManager struct {
	SessionId string
}

func (s *SessionManager) Set(key, value string, expireTime int) {
	cache.HSetKey(s.SessionId, key, value, expireTime)
}

func (s *SessionManager) Get(key string) string {
	value := cache.HGetKey(s.SessionId, key)
	return value
}

func (s *SessionManager) Del() {
	cache.HDelKey(s.SessionId)
}
