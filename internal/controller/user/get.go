package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
)

// UserExists returns whether a user exists in the database
func UserExists(user_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = user_model.GetUser(db, user_id)
	return err == nil
}

// UsersExist returns whether a list of users exists in the database
func UsersExist(user_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsers(db, user_ids)
	return err == nil && len(users) == len(user_ids)
}

// UserExistsFromParams returns whether a user exists in the database
func UserExistsFromName(username string) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{
		"Pseudo": username,
	})
	return err == nil && len(users) > 0
}

// GetUser returns a user from the database
// Selects the user with the given user_id
func GetUser(user_id int) (db_model.User, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return db_model.User{}, err
	}
	defer db_model.CloseDB(db)
	return user_model.GetUser(db, user_id)
}

// GetUsers returns a list of users from the database
// Selects the users with the given user_ids
func GetUsers(user_ids []int) ([]db_model.User, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	return user_model.GetUsers(db, user_ids)
}
