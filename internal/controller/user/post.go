package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	"BeatBoxBox/pkg/utils"
	"errors"
)

func PostUser(username string, email string, raw_password string) error {
	if !utils.CheckPseudoValidity(username) {
		return errors.New("pseudo is invalid, must be between 3 and 32 characters")
	}
	if UserExistsFromName(username) {
		return errors.New("pseudo already exists")
	}
	if !utils.CheckEmailValidity(email) {
		return errors.New("email is invalid, must be less than 256 characters and match the regex pattern")
	}
	if !utils.CheckRawPasswordValidity(raw_password) {
		return errors.New("password is invalid, must be between 6 and 32 characters, contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	hashed_password, err := utils.HashPassword(raw_password)
	if err != nil {
		return err
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	return user_model.CreateUser(db, username, email, hashed_password)
}
