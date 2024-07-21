package role_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

func UpdateRole(db *gorm.DB, role *db_tables.Role, update_map map[string]interface{}) error {
	return db.Model(role).Updates(update_map).Error
}
