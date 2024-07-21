package user_controller

import (
	db_tables "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
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

// GetUserJSON returns a user from the database
func GetUserJSON(user_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return nil, err
	}

	return ConvertUserToJSON(&user)
}

// GetUsersJSON returns a list of users from the database
func GetUsersJSON(user_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	users, err := user_model.GetUsers(db, user_ids)
	if err != nil {
		return nil, err
	} else if users == nil || len(users) != len(user_ids) {
		return nil, custom_errors.NewNotFoundError("some users were not found")
	}
	users_ptr := make([]*db_tables.User, len(users))
	for i, user := range users {
		users_ptr[i] = &user
	}
	return ConvertUsersToJSON(users_ptr)
}

func GetUserIdFromUsername(username string) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)
	users := user_model.GetUsersFromFilters(db, map[string]interface{}{"pseudo": username})
	if users == nil || len(users) == 0 {
		return -1, custom_errors.NewNotFoundError("user not found")
	}
	return users[0].Id, nil
}
