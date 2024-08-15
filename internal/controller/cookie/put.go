package cookie_controller

import (
	db_tables "BeatBoxBox/internal/model"
	cookie_model "BeatBoxBox/internal/model/cookie"
	db_model "BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

func updateAuthTokenIfNearExpiry(auth_cookie *db_tables.AuthCookie) (string, error) {
	if auth_utils.CheckExpiryTimeNear(auth_cookie.ExpirationDate) {
		// Generate a new token
		new_token, new_hash, err := auth_utils.GenerateRandomTokenWithHash()
		if err != nil {
			return "", httputils.NewInternalServerError("failed to generate new token")
		}

		// Update the token in the database
		db, err := db_model.OpenDB()
		if err != nil {
			return "", err
		}
		defer db_model.CloseDB(db)
		err = cookie_model.UpdateCookieAuthToken(db, auth_cookie, new_hash)
		if err != nil {
			return "", err
		}
		return new_token, nil
	}
	return "", nil
}
