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

	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal(err)
	}

}
