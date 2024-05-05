package user_handler

import (
	user_controller "BeatBoxBox/internal/controller/user"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Get the users IDs from the URL
	user_ids_requested := r.URL.Query().Get("user_ids")
	if user_ids_requested == "" {
		http.Error(w, "No user IDs provided, please use user_ids request parameter", http.StatusBadRequest)
		return
	}

	user_ids_str := strings.Split(user_ids_requested, ",")
	user_ids := []int{}
	for _, user_id_str := range user_ids_str {
		user_id, err := strconv.Atoi(user_id_str)
		if err != nil {
			http.Error(w, "Invalid user ID provided, please use a valid positive integer user ID", http.StatusBadRequest)
			return
		}
		user_ids = append(user_ids, user_id)
	}
	users, err := user_controller.GetUsers(user_ids)
	if err != nil {
		http.Error(w, "Internal Server Error when getting all users", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// getUserHandler
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL
	user_id_str := mux.Vars(r)["user_id"]
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		http.Error(w, "Invalid user ID provided, please use a valid integer user ID", http.StatusBadRequest)
		return
	}

	// Get user
	user, err := user_controller.GetUser(user_id)
	if err != nil {
		http.Error(w, "Internal Server Error when getting user", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
