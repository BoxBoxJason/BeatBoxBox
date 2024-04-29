package user_model

import "gorm.io/gorm"

// PostUser creates a new user in the database
func CreateUser(db *gorm.DB, pseudo string, email string, hashed_password string) error {
	new_user := User{
		Pseudo:          pseudo,
		Email:           email,
		Hashed_password: hashed_password,
	}
	return db.Create(&new_user).Error
}
