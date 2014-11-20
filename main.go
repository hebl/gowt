package main

import (
	"github.com/gorilla/mux"
	. "github.com/hebl/gowt/core"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", FilterHandler(Home))
	r.HandleFunc("/login", FilterHandler(Login))
	r.HandleFunc("/logout", FilterHandler(Logout))
	r.HandleFunc("/uc", FilterHandler(UserCenter))
	r.HandleFunc("/reg", FilterHandler(Register))

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", r)

	if err != nil {
		log.Fatal(err)
	}

}
