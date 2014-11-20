package core

import (
	"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

/**
 * 过滤器，检查是否用户登录，记录到Session中
 */
func FilterHandler(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := GetSession(r)

		user := sess.Values[UserSId]

		if user == nil { //用户未登录
			Login(w, r)
		} else {
			f(w, r)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	sess, _ := GetSession(r)

	user := sess.Values[UserSId]
	fmt.Fprintln(w, user)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {
	sess, _ := GetSession(r)
	DeleteSession(r, w, sess)
}

func UserCenter(w http.ResponseWriter, r *http.Request) {

}
