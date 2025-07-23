package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Dragon struct{
	Name string `json:"nome"`
	Classe string `json:"classe"`
	Nivel int `json:"nivel"`
}
func main() {
	dragao := Dragon{Name: "Ryujin",Classe: "Celestial",Nivel:99 }
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Azazel acordou")
	})
	http.HandleFunc("/status",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Status: canal magico estavel")
	})
	http.HandleFunc("/dragon",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Ryujin ruge ao longe...")
	})

	http.HandleFunc("/dragon/info",func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		err :=json.NewEncoder(w).Encode(dragao)
		if err != nil{
			http.Error(w,"erro ao invocar o dragão", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":7070",nil)
}