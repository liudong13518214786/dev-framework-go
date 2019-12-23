package session

import (
	"dev-framework-go/conf"
	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/wonderivan/logger"
	"net/http"
)

type session struct {
	name    string
	request *http.Request
	store   sessions.Store
	session *sessions.Session
	written bool
	writer  http.ResponseWriter
}

//验证登录中间件
func SessionV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := Default(c)
		userUuid := sess.Get("useruuid")
		//userUuid := s.SessionGet(c, "useruuid")
		//cookieManger := util.CookieManger{c, conf.SESSION_NAME, conf.COOKIE_EXPIRE_TIME}
		//cookieId := cookieManger.GetSessionid()
		//sessionManage := util.SessionManager{cookieId}
		//userUuid := session.Get("useruuid")
		code := conf.SUCCESS
		if userUuid == nil {
			code = conf.NOT_LOGIN
		}
		if code != conf.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  conf.GetMessage(code),
				"data": nil,
			})
			c.Abort()
		}
		c.Next()
	}
}

func SessionV2(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		store, _ := redistore.NewRediStore(10, conf.REDIS_NETWORK, conf.REDIS_ADDRESS, conf.REDIS_PASS, []byte(conf.COOKIE_SECRET))
		session := &session{name, c.Request, store, nil, false, c.Writer}
		c.Set(conf.GLOBAL_SESSION, session)
		defer context.Clear(c.Request)
		c.Next()
	}
}

type Session interface {
	Get(key interface{}) interface{}
	Set(key, value interface{})
	Delete(key interface{})
	Clear()
	Save() error
	Options(options Options)
}

type Options struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

func (s *session) Session() *sessions.Session {
	if s.session == nil {
		var err error
		s.session, err = s.store.Get(s.request, s.name)
		if err != nil {
			logger.Error("SESSION ERR:", err)
		}
	}
	return s.session
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key, value interface{}) {
	s.Session().Values[key] = value
	s.written = true
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
	s.written = true
}

func (s *session) Clear() {
	for k := range s.Session().Values {
		s.Delete(k)
	}
}

func (s *session) Options(options Options) {
	s.Session().Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

func (s *session) Save() error {
	if s.written == true {
		err := s.Session().Save(s.request, s.writer)
		if err == nil {
			s.written = false
		}
		return err
	}
	return nil
}

func Default(c *gin.Context) Session {
	return c.MustGet(conf.GLOBAL_SESSION).(Session)
}
