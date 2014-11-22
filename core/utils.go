package core

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RequestInspector(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s, method: %s, vars: %s", r.URL, r.Method, mux.Vars(r))
}

func GenerateCsrfToken() string {
	token := make([]byte, 64)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}
