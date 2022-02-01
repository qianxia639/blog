package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	// 初始化一个cookie储存对象
	// 参数1用户验证，参数2用户加密
	store     = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
	sessionId = "RUOYU_SESSION_ID"
)

// 设置session
func SetSession(ctx *gin.Context, key, value interface{}) error {
	session, err := get(ctx)
	if err != nil {
		return err
	}
	session.Values[key.(string)] = value
	// 30分钟后cookie将会被删除
	// session.Options.MaxAge = 30 * 60
	// return session.Store().Save(ctx.Request,ctx.Writer,session)
	return session.Save(ctx.Request, ctx.Writer)
}

// 获取session
func GetSession(ctx *gin.Context, key interface{}) (interface{}, error) {
	session, err := get(ctx)
	if err != nil {
		return nil, err
	}
	return session.Values[key], nil
}

// 删除session
func RemoveSession(ctx *gin.Context) error {
	session, err := get(ctx)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return sessions.Save(ctx.Request, ctx.Writer)
}

func get(ctx *gin.Context) (*sessions.Session, error) {
	return store.Get(ctx.Request, sessionId)
}
