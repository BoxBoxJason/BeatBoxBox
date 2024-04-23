package user_model

import "gorm.io/gorm"

// PostUser creates a new user in the database
func PostUser(db *gorm.DB, pseudo string, hashed_password string, salt string) error {
	new_user := User{
		Pseudo:          pseudo,
		Hashed_password: hashed_password,
		Salt:            salt,
	}
	return db.Create(&new_user).Error
}
