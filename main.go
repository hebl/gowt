package main

import (
	"github.com/gorilla/mux"
	. "github.com/hebl/Go-Web-Template/core"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", r)

	if err != nil {
		log.Fatal(err)
	}

}
