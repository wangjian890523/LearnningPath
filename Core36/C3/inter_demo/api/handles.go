package api

import (
	"encoding/json"
	"net/http"

	"github.com/wangjian890523/LearnningPath/Core36/C3/inter_demo/db"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.FetchUsers()
	if err != nil {
		http.Error(w, "Failed to fetch  users.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
