package user_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostUser creates a new user in the database
func CreateUser(db *gorm.DB, pseudo string, email string, hashed_password string) error {
	new_user := db_model.User{
		Pseudo:          pseudo,
		Email:           email,
		Hashed_password: hashed_password,
	}
	return db.Create(&new_user).Error
}
