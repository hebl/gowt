package core

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func RequestInspector(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s, method: %s, vars: %s", r.URL, r.Method, mux.Vars(r))
}

func GenerateCsrfToken() string {
	data := string((time.Now().UnixNano() / 1000) % 100000)
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
