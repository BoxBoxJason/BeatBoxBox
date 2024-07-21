package role_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

func RoleExistsFromName(db *gorm.DB, name string) bool {
	var role db_tables.Role
	err := db.Where("name = ?", name).First(&role).Error
	return err == nil
}

func RolesExistFromNames(db *gorm.DB, names []string) bool {
	var roles []db_tables.Role
	err := db.Where("name IN ?", names).Find(&roles).Error
	return err == nil && len(roles) == len(names)
}

func GetRoleByName(db *gorm.DB, name string) (db_tables.Role, error) {
	var role db_tables.Role
	err := db.Where("name = ?", name).First(&role).Error
	return role, err
}

func GetRolesByName(db *gorm.DB, names []string) ([]db_tables.Role, error) {
	var roles []db_tables.Role
	err := db.Where("name IN ?", names).Find(&roles).Error
	return roles, err
}

func GetRoleIdByName(db *gorm.DB, name string) (int, error) {
	role, err := GetRoleByName(db, name)
	if err != nil {
		return -1, err
	}
	return role.Id, err
}
