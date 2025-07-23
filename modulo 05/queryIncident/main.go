package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type package15 struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Weapon string `json:"weapon"`
	Level  int    `json:"level"`
}

func main() {
	azazel := package15{Name: "Azazel", Class: "Mage", Weapon: "the negative", Level: 65}
	kaya := package15{Name: "Kaya", Class: "Sniper", Weapon: "Fal", Level: 35}
	marcos := package15{Name: "Marcos", Class: "Velocist", Weapon: "Katana", Level: 65}

	theFuckingPackage15:= []package15{azazel,kaya,marcos}
	http.HandleFunc("/package",func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		err:= json.NewEncoder(w).Encode(theFuckingPackage15)
		if err != nil{
			http.Error(w,"they are not here", http.StatusInternalServerError)
		}
	
	})


	http.HandleFunc("/buscar",func(w http.ResponseWriter, r *http.Request) {
		classe := r.URL.Query().Get("class")
		nivel := r.URL.Query().Get("level")

		if classe == "" || nivel == "" {
			http.Error(w, "Faltam parâmetros: classe ou nivel", http.StatusBadRequest)
			return
		}
		var selected []package15
		for _, p:= range theFuckingPackage15{
			if p.Class == classe && strconv.Itoa(p.Level) == nivel{
				selected = append(selected,p )
			}
		}
		err:= json.NewEncoder(w).Encode(selected)
		if err != nil{
			http.Error(w,"its not here", http.StatusInternalServerError)
		}
		
	})

	http.ListenAndServe(":7070",nil)
}