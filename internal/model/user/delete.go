package user_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// DeleteUser deletes an existing user from the database
func DeleteUser(db *gorm.DB, user_id int) error {
	return db.Delete(&db_model.User{}, user_id).Error
}

// DeleteUsers deletes existing users from the database
func DeleteUsers(db *gorm.DB, user_ids []int) error {
	return db.Where("Id IN ?", user_ids).Delete(&db_model.User{}).Error
}
