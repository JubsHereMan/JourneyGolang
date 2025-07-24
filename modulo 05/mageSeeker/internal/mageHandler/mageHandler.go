package magehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"mageSeeker/internal/MageStruct"
	"github.com/gorilla/mux"
)

func MageHandler(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	mage:= vars["name"]
	fmt.Fprintf(w,"você invocou o mago %s",mage)
	w.Header().Set("Content-Type","application/json")
	var selected []magestruct.Mage
	for _, m:= range magestruct.Mages{
		if m.Name == mage{
			selected = append(selected, m)
		}
	}
	err := json.NewEncoder(w).Encode(selected)
	if err != nil{
		http.Error(w,"error", http.StatusInternalServerError)
	}
}