package internalPost

import (
	"allMembers/members"
	"encoding/json"
	"net/http"
)
func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var newChar members.Character

	err := json.NewDecoder(r.Body).Decode(&newChar)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	members.Members.Members = append(members.Members.Members, newChar)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newChar)
}
