package user_handler

import (
	user_controller "BeatBoxBox/internal/controller/user"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getUserHandler gets a user by their ID
// @Summary Get a user by their ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} User "OK"
// @Failure 400 {string} string "Invalid user ID provided"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/users/{user_id} [get]
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	user_ids_str := r.URL.Query().Get("users_ids")
	users_ids, err := format_utils.ConvertStringToIntArray(user_ids_str, ",")
	if err != nil {
		http.Error(w, "No user IDs provided, please use user_ids request parameter", http.StatusBadRequest)
		return
	}
	if !user_controller.UsersExist(users_ids) {
		http.Error(w, "One or more users do not exist", http.StatusNotFound)
		return
	}

	users, err := user_controller.GetUsers(users_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// getUserHandler gets a user by their ID
// @Summary Get a user by their ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} User "OK"
// @Failure 400 {string} string "Invalid user ID provided"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/users/{user_id} [get]
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL
	user_id_str := mux.Vars(r)["user_id"]
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		http.Error(w, "Invalid user ID provided, please use a valid integer user ID", http.StatusBadRequest)
		return
	}
	if !user_controller.UserExists(user_id) {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	// Get user
	user, err := user_controller.GetUser(user_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(user)
}
