package internal

import (
	"allMembers/members"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetByClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	class := params["class"]

	for _, c := range members.Members.Members {
		if c.Class == class {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "character not found", http.StatusNotFound)
}
