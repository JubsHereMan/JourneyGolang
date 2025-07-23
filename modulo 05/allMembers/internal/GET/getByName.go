package internal

import (
	"allMembers/members"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	for _, c := range members.Members.Members {
		if c.Name == name {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "character not found", http.StatusNotFound)
}
