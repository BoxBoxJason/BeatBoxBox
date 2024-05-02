package user_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdateUser updates an existing user in the database
func UpdateUser(db *gorm.DB, user_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.User{}).Where("Id = ?", user_id).Updates(update_map).Error
}
