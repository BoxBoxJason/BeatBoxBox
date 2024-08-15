package user_controller

import (
	db_tables "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	"BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

// DeleteUser deletes a user by its id
func DeleteUser(user_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return err
	}
	return user_model.DeleteUserFromRecord(db, &user)
}

// DeleteUsers deletes users by their ids
func DeleteUsers(user_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsers(db, user_ids)
	if err != nil {
		return err
	} else if users == nil || len(users) != len(user_ids) {
		return httputils.NewNotFoundError("some users were not found")
	}
	users_ptr := make([]*db_tables.User, len(users))
	for i, user := range users {
		users_ptr[i] = &user
	}
	return user_model.DeleteUsersFromRecords(db, users_ptr)
}
