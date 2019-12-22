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
	SessionInit(sid string) (Session, error) //session初始化
	SessionRead(sid string) (Session, error) //session读取（如果读取不到，则创建一个新的）
	SessionDestroy(sid string) error         // session的销毁（这个是用户主动销毁）
	SessionGC(maxLifeTime int64)             //session的回收（根据超时时间来回收）
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
