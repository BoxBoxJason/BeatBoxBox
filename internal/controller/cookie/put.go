package cookie_controller

import (
	db_model "BeatBoxBox/internal/model"
	cookie_model "BeatBoxBox/internal/model/cookie"
	"time"
)

func updateAuthTokenIfNearExpiry(auth_cookie db_model.AuthCookie) (string, error) {
	if checkTokenNearExpiry(auth_cookie.ExpirationDate) {
		// Generate a new token
		new_token, new_hash, err := generateRandomTokenWithHash()
		if err != nil {
			return "", err
		}

		// Update the token in the database
		db, err := db_model.OpenDB()
		if err != nil {
			return "", err
		}
		defer db_model.CloseDB(db)
		cookie_model.UpdateCookieAuthToken(db, auth_cookie.Id, new_hash)
		return new_token, nil
	}
	return "", nil
}

func checkTokenNearExpiry(expiry_time int64) bool {
	remaining_time := expiry_time - time.Now().Unix()
	return remaining_time < int64(cookie_model.DEFAULT_TOKEN_EXPIRATION.Seconds())
}
