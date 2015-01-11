package core

/**
 * 系统程序类
 *
 */

import (
	//"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

type WTApp struct {
	Router *mux.Router
}

type WTContext struct {
	user    User
	session *sessions.Session
	res     http.ResponseWriter
	rep     *http.Request
}

func NewWTApp() *WTApp {
	app := &WTApp{
		Router: mux.NewRouter(),
	}

	return app
}

// func (app *WTApp) SetupRoute() {
// 	// 继承使用，设置路由和
// }

func (app *WTApp) SetupRoute() {
	app.Router.HandleFunc("/", FilterHandler(Home))
	app.Router.HandleFunc("/login", Login)
	app.Router.HandleFunc("/logout", FilterHandler(Logout))
	app.Router.HandleFunc("/uc", FilterHandler(UserCenter))
	app.Router.HandleFunc("/reg", FilterHandler(Register))
}

func (app *WTApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, _ := GetSession(r)
	csrf_token, ok := sess.Values["csrf_token"]

	if !ok {
		csrf_token = GenerateCsrfToken()
		sess.Values["csrf_token"] = csrf_token
		sess.Save(r, w)
	}

	context.Set(r, "csrf_token", csrf_token)

	if r.Method == "POST" && r.FormValue("csrf_token") != csrf_token {
		http.Error(w, "Fobidden", http.StatusForbidden)
	} else {
		app.Router.ServeHTTP(w, r)
	}

}

// 改编自 https://github.com/achun/typepress/blob/master/src/global/global.go
// 获取Session
func GetSession(r *http.Request) (*sessions.Session, error) {
	sess, err := SessionStore.Get(r, SessionName)
	if err != nil { // 如果无，则新建Session Cookie
		sess, err = NewSession(r)
	}
	return sess, err
}

// NewSession
func NewSession(r *http.Request) (*sessions.Session, error) {
	sess, err := SessionStore.New(r, SessionName)
	if err != nil {
		sess, err = SessionStore.New(r, SessionName)
	}
	sess.Options.HttpOnly = true
	sess.Options.MaxAge = 86400 * 14 // 两周

	return sess, err
}

// SaveSession
func SaveSession(r *http.Request, w http.ResponseWriter, sess *sessions.Session) error {
	return sess.Save(r, w)
}

// DeleteSession
func DeleteSession(r *http.Request, w http.ResponseWriter, sess *sessions.Session) error {
	sess.Options.MaxAge = -1

	return sess.Save(r, w)
}
