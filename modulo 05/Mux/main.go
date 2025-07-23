package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/summon/{creature}",summonHandler)
	http.ListenAndServe(":7070",r)
}

func summonHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	creature := vars["creature"]
	fmt.Fprintf(w, "você invocou a criatura: %s", creature)

}

