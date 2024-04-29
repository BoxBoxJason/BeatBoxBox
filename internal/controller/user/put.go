package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
)

func UpdateUser(user_id int, user_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return user_model.UpdateUser(db, user_id, user_map)
}
