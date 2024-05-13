package cookie_controller

import (
	db_model "BeatBoxBox/internal/model"
	cookie_model "BeatBoxBox/internal/model/cookie"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"crypto/rand"
	"encoding/hex"
)

// generateToken generates a random token of 128 bits
func generateRandomTokenWithHash() (string, string, error) {
	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", err
	}
	raw_token := hex.EncodeToString(bytes)
	hashed_token, err := auth_utils.HashString(raw_token)
	if err != nil {
		return "", "", err
	}
	return raw_token, hashed_token, nil
}

func PostAuthToken(user_id int) (string, error) {
	// Generate a random token and hash it
	raw_auth_token, hashed_auth_token, err := generateRandomTokenWithHash()
	if err != nil {
		return "", err
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
