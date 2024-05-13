package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	custom_errors "BeatBoxBox/pkg/errors"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
)

func PostUser(username string, email string, raw_password string) (int, error) {
	if !format_utils.CheckPseudoValidity(username) {
		return -1, custom_errors.NewBadRequestError("pseudo is invalid, must be between 3 and 32 characters")
	}
	if UserExistsFromName(username) {
		return -1, custom_errors.NewBadRequestError("pseudo already exists")
	}
	if !format_utils.CheckEmailValidity(email) {
		return -1, custom_errors.NewBadRequestError("email is invalid, must be less than 256 characters and match the regex pattern")
	}
	if !format_utils.CheckRawPasswordValidity(raw_password) {
		return -1, custom_errors.NewBadRequestError("password is invalid, must be between 6 and 32 characters, contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	hashed_password, err := auth_utils.HashString(raw_password)
	if err != nil {
		return -1, err
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return user_model.CreateUser(db, username, email, hashed_password)
}
