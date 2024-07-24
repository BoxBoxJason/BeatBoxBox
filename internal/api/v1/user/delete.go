package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// deleteUserHandler deletes a user with the given ID
// @Summary Delete a user by their ID
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid user ID provided, please use a valid integer user ID"
// @Failure 404 {string} string "User does not exist"
// @Failure 500 {string} string "Internal server error when deleting user"
// @Router /api/users/{user_id} [delete]
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL
	user_id_str := mux.Vars(r)["user_id"]
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		http.Error(w, "Invalid user ID provided, please use a valid integer user ID", http.StatusBadRequest)
		return
	}

	// Check if user exists
	if !user_controller.UserExists(user_id) {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	// Delete user
	err = user_controller.DeleteUser(user_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
}

// deleteUsersHandler deletes users with the given IDs
// @Summary Delete users by their IDs
// @Description Delete users by their IDs
// @Tags users
// @Accept json
// @Produce json
// @Param users_ids query []int true "User IDs"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "No user IDs provided, please use user_ids request parameter"
// @Failure 404 {string} string "One or more users do not exist"
// @Failure 500 {string} string "Internal server error when deleting users"
// @Router /api/users [delete]
func deleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	users_ids_str := r.URL.Query().Get("users_ids")
	users_ids, err := format_utils.ConvertStringToIntArray(users_ids_str, ",")
	if err != nil {
		http.Error(w, "Invalid user IDs provided, please use a valid integer user IDs", http.StatusBadRequest)
		return
	}
	if !user_controller.UsersExist(users_ids) {
		http.Error(w, "One or more users do not exist", http.StatusNotFound)
		return
	}

	err = user_controller.DeleteUsers(users_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
