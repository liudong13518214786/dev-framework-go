package session

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
)

//定义session
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionId() string
}

//session管理器
type SessionManager interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

func SessionGet(c *gin.Context, key string) string {
	cookieManger := &util.CookieManger{C: c, SessionName: conf.SESSION_NAME, ExpireTime: conf.COOKIE_EXPIRE_TIME}
	cookieId := cookieManger.GetSessionid()
	sessionManage := &util.SessionManager{SessionId: cookieId}
	return sessionManage.Get(key)
}

func SessionSet(c *gin.Context, key string, value string, expireTime int) bool {
	cookieManger := &util.CookieManger{C: c, SessionName: conf.SESSION_NAME, ExpireTime: conf.COOKIE_EXPIRE_TIME}
	cookieId := cookieManger.GetSessionid()
	sessionManage := &util.SessionManager{SessionId: cookieId}
	return sessionManage.Set(key, value, expireTime)
}

func SessionDel(c *gin.Context) {
	cookieManger := &util.CookieManger{C: c, SessionName: conf.SESSION_NAME, ExpireTime: conf.COOKIE_EXPIRE_TIME}
	cookieId := cookieManger.GetSessionid()
	sessionManage := &util.SessionManager{SessionId: cookieId}
	sessionManage.Del()
}
