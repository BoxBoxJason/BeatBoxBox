package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	"BeatBoxBox/pkg/utils"
	"errors"
)

func AttemptLogin(username_or_email string, raw_password string) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{"email": username_or_email})
	if err != nil {
		users, err = user_model.GetUsersFromFilters(db, map[string]interface{}{"pseudo": username_or_email})
	}
	if err != nil || len(users) == 0 {
		return errors.New("user not found")
	}
	user := users[0]

	hashed_password := user.Hashed_password
	if !utils.ComparePasswords(hashed_password, raw_password) {
		return errors.New("wrong password")
	}
	return nil
}
