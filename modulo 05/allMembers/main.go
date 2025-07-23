package main

import (
	"allMembers/internal/GET"
	"allMembers/internal/POST"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/character/{name}", internal.GetByName).Methods("GET")
	router.HandleFunc("/class/{class}", internal.GetByClass).Methods("GET")
	router.HandleFunc("/character", internalPost.CreateCharacter).Methods("POST")
	http.ListenAndServe(":7070",router)
}
