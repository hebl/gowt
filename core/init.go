package core

/**
 * 全局变量，常量
 *
 *
 */

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"gopkg.in/pg.v3"
	//"net/http"
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
	UserSId     = "UserId"

	DB_User = "hebl"
	DB_Name = "test2"
)

// 初始化

func init() {
	db = pg.Connect(&pg.Options{
		User:     DB_User,
		Database: DB_Name,
	})

	SessionStore = sessions.NewCookieStore([]byte(sessionSeed))
	gob.Register(&User{}) //注册User模型

	//
	TplPath = "templates/layout.html"
}
