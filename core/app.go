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
	Router       *mux.Router
	SessionStore *sessions.CookieStore
}

func NewWTApp() *WTApp {
	app := &WTApp{
		Router:       mux.NewRouter(),
		SessionStore: sessions.NewCookieStore([]byte(sessionSeed)),
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
	sess, _ := app.GetSession(r)
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
func (app *WTApp) GetSession(r *http.Request) (*sessions.Session, error) {
	sess, err := app.SessionStore.Get(r, SessionName)
	if err != nil { // 如果无，则新建Session Cookie
		sess, err = app.NewSession(r)
	}
	return sess, err
}

// NewSession
func (app *WTApp) NewSession(r *http.Request) (*sessions.Session, error) {
	sess, err := app.SessionStore.New(r, SessionName)
	if err != nil {
		sess, err = app.SessionStore.New(r, SessionName)
	}
	sess.Options.HttpOnly = true
	sess.Options.MaxAge = 86400 * 14 // 两周

	return sess, err
}

// SaveSession
func (app *WTApp) SaveSession(r *http.Request, w http.ResponseWriter, sess *sessions.Session) error {
	return sess.Save(r, w)
}

// DeleteSession
func (app *WTApp) DeleteSession(r *http.Request, w http.ResponseWriter, sess *sessions.Session) error {
	sess.Options.MaxAge = -1

	return sess.Save(r, w)
}
