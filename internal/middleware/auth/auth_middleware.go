package auth_middleware

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	cookie_model "BeatBoxBox/internal/model/cookie"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve cookie & jwt info
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusUnauthorized)
			return
		}
		user_id, auth_token, err := auth_utils.ParseAuthJWT(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusUnauthorized)
			return
		}

		// Check if auth token matches & pass the new token if one was generated
		auth_token_matches, new_token, err := cookie_controller.CheckAuthTokenMatches(user_id, auth_token)
		if err != nil || !auth_token_matches {
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
				Expires:  time.Now().Add(cookie_model.DEFAULT_TOKEN_EXPIRATION),
			})
		}

		next.ServeHTTP(w, r)
	})
}
