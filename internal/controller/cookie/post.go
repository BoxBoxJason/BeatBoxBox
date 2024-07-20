package cookie_controller

import (
	cookie_model "BeatBoxBox/internal/model/cookie"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
)

func PostAuthToken(user_id int) (string, error) {
	// Generate a random token and hash it
	raw_auth_token, hashed_auth_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		return "", custom_errors.NewInternalServerError("failed to generate auth token")
	}

	// Store the hashed token in the database
	db, err := db_model.OpenDB()
	if err != nil {
		return "", err
	}
	defer db_model.CloseDB(db)
	_, err = cookie_model.CreateCookie(db, hashed_auth_token, user_id)
	if err != nil {
		return "", err
	}

	return raw_auth_token, nil
}
