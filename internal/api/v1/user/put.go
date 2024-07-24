package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// putUserHandler updates a user by their ID
// @Summary Update a user by their ID
// @Description Update a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param username formData string false "Username"
// @Param email formData string false "Email"
// @Param illustration formData file false "Illustration"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid user ID provided"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/users/{user_id} [put]
func putUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
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

	// Parse the url parameters and retrieve only authorized ones
	update_dict := make(map[string]interface{})
	username := r.FormValue("username")
	if username != "" {
		update_dict["pseudo"] = username
	}
	email := r.FormValue("email")
	if email != "" {
		update_dict["email"] = email
	}

	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "users")
		if err != nil || illustration_file_name == file_utils.DEFAULT_ILLUSTRATION_FILE {
			custom_errors.SendErrorToClient(err, w, "")
			return
		}
		update_dict["illustration"] = illustration_file_name
	}
	err = user_controller.UpdateUser(user_id, update_dict)
	if err != nil {
		http.Error(w, "Error when updating user: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
