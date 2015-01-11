package core

import (
	"fmt"
	"github.com/gorilla/context"
	//"github.com/gorilla/sessions"
	"net/http"
	//"reflect"
)

type MyHandleFunc func(ctx *WTContext) http.Handler

/**
 * 过滤器，检查是否用户登录，记录到Session中
 */
func FilterHandler(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := GetSession(r)

		user := sess.Values[UserSId]

		if user == nil { //用户未登录
			http.Redirect(w, r, "/login", 302)
		} else {
			f(w, r)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	sess, _ := GetSession(r)

	user := sess.Values[UserSId]
	fmt.Println(user)

	RenderHtml(w, "index", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx := make(map[string]interface{})
		ctx["csrf_token"] = context.Get(r, "csrf_token")
		RenderHtml(w, "login", ctx)
	} else {
		fmt.Println("OK")
		RenderHtml(w, "index", nil)
	}

}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
	RenderHtml(w, "index", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	RenderHtml(w, "reg", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//sess, _ := GetSession(r)
	//DeleteSession(r, w, sess)
}

func UserCenter(w http.ResponseWriter, r *http.Request) {
	RenderHtml(w, "uc", nil)
}

// 404
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	RenderHtml(w, "404", nil)
}
