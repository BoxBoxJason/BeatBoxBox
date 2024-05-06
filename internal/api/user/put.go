package user_handler

import (
	user_controller "BeatBoxBox/internal/controller/user"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func putUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL
	user_id_str := mux.Vars(r)["user_id"]
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		http.Error(w, "Invalid user ID provided, please use a valid integer user ID", http.StatusBadRequest)
		return
	}

	// Parse the url parameters and retrieve only authorized ones
	update_dict, err := parseURLParams(r.URL.Query())
	if err != nil {
		http.Error(w, "Invalid URL parameters provided: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = user_controller.UpdateUser(user_id, update_dict)
	if err != nil {
		http.Error(w, "Error when updating user: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
