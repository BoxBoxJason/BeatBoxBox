package user_handler

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	user_controller "BeatBoxBox/internal/controller/user"
	cookie_model "BeatBoxBox/internal/model/cookie"
	custom_errors "BeatBoxBox/pkg/errors"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"net/http"
	"time"
)

// loginHandler logs in the user with the given username or email and password
// @Summary Log in the user
// @Description Log in the user
// @Tags users
// @Accept json
// @Produce json
// @Param username_or_email formData string true "Username or Email"
// @Param password formData string true "Password"
// @Success 200 {string} string "OK"
// @Failure 401 {string} string "Invalid username / password"
// @Router /api/users/login [post]
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	attempt_username_or_email := r.PostFormValue("username_or_email")
	raw_password := r.PostFormValue("password")
	if len(attempt_username_or_email) < 3 || len(raw_password) < 6 {
		http.Error(w, "Invalid Username / Password", http.StatusUnauthorized)
		return
	}

	user_id, session_token, err := user_controller.AttemptLogin(attempt_username_or_email, raw_password)
	if err != nil {
		http.Error(w, "Invalid username / password", http.StatusUnauthorized)
		return
	}

	updateSessionCookie(w, user_id, session_token)

	http.Redirect(w, r, "/", http.StatusFound)
}

// logoutHandler logs out the user
// @Summary Log out the user
// @Description Log out the user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Router /api/users/logout [post]
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/auth", http.StatusUnauthorized)
		return
	}
	user_id, token, err := auth_utils.ParseAuthJWT(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	err = cookie_controller.DeleteMatchingAuthToken(user_id, token)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	deleteSessionCookie(w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func updateSessionCookie(w http.ResponseWriter, user_id int, raw_auth_token string) {
	new_jwt, err := auth_utils.CreateAuthJWT(user_id, raw_auth_token)
	if err != nil {
		http.Error(w, "Internal Server Error when creating new auth JWT", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    new_jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(cookie_model.DEFAULT_TOKEN_EXPIRATION),
	})
}

func deleteSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-time.Hour),
	})
	user_id, auth_token, err := auth_utils.ParseAuthJWT("session_token")
	if err != nil {
		http.Error(w, "Error parsing session token", http.StatusUnauthorized)
		return
	}
	go cookie_controller.DeleteMatchingAuthToken(user_id, auth_token) // Don't wait for this to finish
}
