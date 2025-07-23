package main

import (
	"encoding/json"
	"net/http"
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
	http.ListenAndServe(":7070",nil)
}