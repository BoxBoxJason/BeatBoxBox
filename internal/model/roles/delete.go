package role_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

func DeleteRoleFromRecord(db *gorm.DB, role *db_tables.Role) error {
	return db.Delete(role).Error
}

func DeleteRolesFromRecords(db *gorm.DB, roles []*db_tables.Role) error {
	return db.Delete(roles).Error
}
