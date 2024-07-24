package user_handler_v1

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	user_controller "BeatBoxBox/internal/controller/user"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
)

// RegisterHandler handles the registration of a user
// @Summary Register a user
// @Description Register a user
// @Tags User
// @Accept json
// @Produce json
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param password_confirm formData string true "Password confirmation"
// @Success 201 {string} string "User created"
// @Failure 400 {string} string "Invalid username, must be between 3 and 32 characters"
// @Failure 400 {string} string "Invalid email, must be less than 256 characters be a valid email pattern"
// @Failure 400 {string} string "Invalid password, must be at least 6 characters long with at least one special character, one uppercase, one lowercase and one number"
// @Failure 400 {string} string "Passwords do not match"
// @Failure 400 {string} string "Username already exists"
// @Failure 400 {string} string "Email already exists"
// @Failure 500 {string} string "Error creating user"
// @Failure 500 {string} string "Error creating auth token, the account was created successfuly"
// @Router /api/users/register [post]
func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	// Retrieve the form values
	username := r.FormValue("username")
	email := r.FormValue("email")
	raw_password := r.FormValue("password")
	raw_password_confirm := r.FormValue("password_confirm")

	// Check the form values
	if !format_utils.CheckPseudoValidity(username) {
		http.Error(w, "Invalid username, must be between 3 and 32 characters", http.StatusBadRequest)
		return
	}
	if !format_utils.CheckEmailValidity(email) {
		http.Error(w, "Invalid email, must be less than 256 characters be a valid email pattern", http.StatusBadRequest)
		return
	}
	if !format_utils.CheckRawPasswordValidity(raw_password) {
		http.Error(w, "Invalid password, must be at least 6 characters long with at least one special character, one uppercase, one lowercase and one number", http.StatusBadRequest)
		return
	}
	if raw_password != raw_password_confirm {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	if user_controller.UserExistsFromName(username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	if user_controller.UserExistsFromEmail(email) {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	// Create the user
	user_id, err := user_controller.PostUser(username, email, raw_password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	raw_auth_token, err := cookie_controller.PostAuthToken(user_id)
	if err != nil {
		http.Error(w, "Error creating auth token, the account was created successfuly", http.StatusInternalServerError)
		http.Redirect(w, r, "/auth", http.StatusInternalServerError)
		return
	}

	// Set the session token in the cookie
	updateSessionCookie(w, user_id, raw_auth_token)

	http.Redirect(w, r, "/", http.StatusCreated)
}
