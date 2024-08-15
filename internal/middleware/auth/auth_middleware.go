package auth_middleware

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id, new_token := AuthenticateUser(r)
		if user_id < 0 {
			http.Redirect(w, r, "/auth", http.StatusUnauthorized)
			return
		}
		if new_token != "" {
			new_jwt, err := auth_utils.CreateAuthJWT(user_id, new_token)
			if err != nil {
				http.Error(w, "Internal Server Error when creating new auth JWT", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "session_token",
				Value:    new_jwt,
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				Expires:  time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION),
			})
		}
		next.ServeHTTP(w, r)
	})
}

func AuthenticateUser(r *http.Request) (int, string) {
	jwt_token := ""
	// Attempt to retrieve cookie & jwt info
	cookie, err := r.Cookie("session_token")
	if err != nil {
		jwt_token = r.Header.Get("Authorization")
	} else {
		jwt_token = cookie.Value
	}

	user_id, auth_token, err := auth_utils.ParseAuthJWT(jwt_token)
	if err != nil || user_id < 0 || auth_token == "" {
		return -1, ""
	}
	// Check if auth token matches & pass the new token if one was generated
	auth_token_matches, new_token, err := cookie_controller.CheckAuthTokenMatches(user_id, auth_token)
	if err != nil || !auth_token_matches {
		return -1, ""
	}
	return user_id, new_token
}

func HasWritePrivileges(r *http.Request) error { // TODO
	jwt_token := ""
	// Attempt to retrieve cookie & jwt info
	cookie, err := r.Cookie("session_token")
	if err != nil {
		jwt_token = r.Header.Get("Authorization")
	} else {
		jwt_token = cookie.Value
	}
	if jwt_token == "" {
		return httputils.NewUnauthorizedError("No auth JWT token found")
	}
	return nil
}
