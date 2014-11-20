package core

/**
 * 全局变量，常量
 *
 *
 */

import (
	"github.com/gorilla/sessions"
	"gopkg.in/pg.v3"
)

// 全局变量
var (
	db           *pg.DB         // 数据库连接池
	SessionStore sessions.Store // Session
	TplPath      string         // 模板根目录
)

// 全局常量
const (
	SessionName = "JSESSIONID"                                                //模拟JSP名称
	sessionSeed = "UkFCRUFBQVZfLUNBQUVHYzNSeWFXNW5EQXNBQ1hWelpYSmZjMlZ6Y3dvc" // CookieStore Session的初始KeyPairs

	DB_User = "hebl"
	DB_Name = "test2"
)

// 初始化

func init() {
	Db := pg.Connect(&pg.Options{
		User:     DB_User,
		Database: DB_Name,
	})
}

// 改编自 https://github.com/achun/typepress/blob/master/src/global/global.go
// 获取Session
func GetSession(r *http.Request) (*sessions.Session, error) {
	sess, err := SessionStore.Get(r, SessionName)
	if err != nil { // 如果无，则新建Session Cookie
		r.Header.Del("Cookie")
		sess, err = NewSession(r)
	}
	return sess, err
}

// NewSession
func NewSession(r *http.Request) (*sessions.Session, error) {
	sess, err := SessionStore.New(r, SessionName)
	if err != nil {
		r.Header.Del("Cookie")
		sess, err = SessionStore.New(r, SessionName)
	}
	sess.Options.HttpOnly = true
	sess.Options.MaxAge = 86400 * 14 // 两周
	return sess, err
}

// SaveSession
func SaveSession(r *http.Request, wr http.ResponseWriter, sess *sessions.Session) error {
	return sess.Save(r, wr)
}
