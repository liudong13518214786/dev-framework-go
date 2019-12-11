package util

import "C"
import (
	"dev-framework-go/conf"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wonderivan/logger"
	"strings"
	"sync"
)

type CookieManger struct {
	C           *gin.Context
	SessionName string
	ExpireTime  int
	mu          sync.Mutex
}

func (s *CookieManger) SetSecureCookie(sessionId string) {
	secureCookie := createSignedValue(s.SessionName, sessionId)
	s.C.SetCookie(s.SessionName, secureCookie, s.ExpireTime, "/", conf.DOMAIN, false, true)
}

func (s *CookieManger) GetSessionid() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	sessionId := s.GetSecureCookie()
	if sessionId == "" {
		sessionId = CreatSessionId()
		s.SetSecureCookie(sessionId)
	}
	return sessionId
}

func (s *CookieManger) GetSecureCookie() string {
	res, err := s.C.Cookie(s.SessionName)
	if err != nil {
		//没有获取到，为登陆状态
		return ""
	}
	sessionId := decodeSignedValue(res)
	return sessionId
}

func CreatSessionId() string {
	uuix := uuid.New()
	sessionId := strings.Replace(uuix.String(), "-", "", -1)
	return sessionId
}

func decodeSignedValue(sign string) string {
	res, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		logger.Error("[BASE64 DECODE ERROR]sign=>%s", sign)
		return ""
	}
	result := strings.Split(string(res), ":")
	if len(result) != 3 {
		logger.Error("[COOIKE ERROR]300")
		return ""
	}
	if result[2] != conf.COOKIE_SECRET {
		logger.Error("[COOIKE ERROR]301")
	}
	return result[1]
}

func createSignedValue(name, value string) string {
	sign := []byte(name + ":" + value + ":" + conf.COOKIE_SECRET)
	return base64.StdEncoding.EncodeToString(sign)
}
