package user_controller

import (
	cookie_controller "BeatBoxBox/internal/controller/cookie"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"BeatBoxBox/pkg/utils/httputils"
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
	users := user_model.GetUsersFromFilters(db, map[string]interface{}{"email": username_or_email})
	if len(users) == 0 {
		users = user_model.GetUsersFromFilters(db, map[string]interface{}{"pseudo": username_or_email})
	}
	if len(users) == 0 {
		return -1, "", errors.New("user not found")
	}
	user := users[0]

	// Check the password
	if !auth_utils.CompareHash(user.HashedPassword, raw_password) {
		return -1, "", errors.New("wrong password")
	}

	// Generate & Set a session token
	raw_session_token, err := cookie_controller.PostAuthToken(user.Id)
	if err != nil {
		return -1, "", err
	}

	return user.Id, raw_session_token, nil
}

func AttemptLoginFromEmail(email string, raw_password string) (int, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, "", err
	}
	defer db_model.CloseDB(db)
	users := user_model.GetUsersFromFilters(db, map[string]interface{}{"email": email})
	if len(users) == 0 {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	user := users[0]
	if !auth_utils.CompareHash(user.HashedPassword, raw_password) {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	raw_session_token, err := cookie_controller.PostAuthToken(user.Id)
	if err != nil {
		return -1, "", httputils.NewInternalServerError("Internal Server Error when creating new auth JWT")
	}
	return user.Id, raw_session_token, nil
}

func AttemptLoginFromId(id int, raw_password string) (int, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, "", err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, id)
	if err != nil {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	if !auth_utils.CompareHash(user.HashedPassword, raw_password) {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	raw_session_token, err := cookie_controller.PostAuthToken(user.Id)
	if err != nil {
		return -1, "", httputils.NewInternalServerError("Internal Server Error when creating new auth JWT")
	}
	return user.Id, raw_session_token, nil
}

func AttemptLoginFromUsername(username string, raw_password string) (int, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, "", err
	}
	defer db_model.CloseDB(db)
	users := user_model.GetUsersFromFilters(db, map[string]interface{}{"pseudo": username})
	if len(users) == 0 {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	user := users[0]
	if !auth_utils.CompareHash(user.HashedPassword, raw_password) {
		return -1, "", httputils.NewUnauthorizedError("Authentication failed")
	}
	raw_session_token, err := cookie_controller.PostAuthToken(user.Id)
	if err != nil {
		return -1, "", httputils.NewInternalServerError("Internal Server Error when creating new auth JWT")
	}
	return user.Id, raw_session_token, nil
}
