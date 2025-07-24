package main

import (
	"mageSeeker/internal/mageHandler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mage/{name}",magehandler.MageHandler).Methods("GET")
	
	http.ListenAndServe(":7070",r)
}