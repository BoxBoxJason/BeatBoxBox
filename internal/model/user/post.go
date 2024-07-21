package user_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostUser creates a new user in the database
func CreateUser(db *gorm.DB, pseudo string, email string, hashed_password string, avatar_file_name string) (int, error) {
	new_user := db_model.User{
		Pseudo:          pseudo,
		Email:           email,
		Hashed_password: hashed_password,
		Illustration:    avatar_file_name,
	}
	err := db.Create(&new_user).Error
	if err != nil {
		return -1, err
	}
	return new_user.Id, nil
}
