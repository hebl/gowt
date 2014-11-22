package main

import (
	//"github.com/gorilla/mux"
	. "github.com/hebl/gowt/core"
	"log"
	"net/http"
)

func main() {

	wtapp := NewWTApp()
	wtapp.SetupRoute()

	http.Handle("/", wtapp)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal(err)
	}

}
