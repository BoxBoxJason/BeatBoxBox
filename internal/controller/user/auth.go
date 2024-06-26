package user_controller

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"errors"
)

// AttemptLogin attempts to login a user with the given username or email and password
func AttemptLogin(username_or_email string, raw_password string) (int, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, "", err
	}
	defer db_model.CloseDB(db)

	// Get the user from the database
	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{"Email": username_or_email})
	if err != nil {
		users, err = user_model.GetUsersFromFilters(db, map[string]interface{}{"Pseudo": username_or_email})
	}
	if err != nil || len(users) == 0 {
		return -1, "", errors.New("user not found")
	}
	user := users[0]

	// Check the password
	if !auth_utils.CompareHash(user.Hashed_password, raw_password) {
		return -1, "", errors.New("wrong password")
	}

	// Generate & Set a session token
	raw_session_token, err := cookie_controller.PostAuthToken(user.Id)
	if err != nil {
		return -1, "", err
	}

	return user.Id, raw_session_token, nil
}
