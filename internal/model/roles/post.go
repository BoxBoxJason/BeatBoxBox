package role_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	"gorm.io/gorm"
)

func init() {
	db, err := db_model.OpenDB()
	if err != nil {
		logger.Error("Failed to connect database: ", err)
	}
	defer db_model.CloseDB(db)
	ROLES_NAMES := []string{"admin", "user", "artist", "moderator"}
	if RolesExistFromNames(db, ROLES_NAMES) {
		return
	}
	ROLES := []map[string]string{
		{"name": "admin", "description": "BeatBoxBox administrator, can do anything"},
		{"name": "user", "description": "BeatBoxBox user, can listen to music and create playlists"},
		{"name": "artist", "description": "BeatBoxBox artist, can upload musics and albums"},
		{"name": "moderator", "description": "BeatBoxBox moderator, can upload and edit musics and albums"},
	}
	for _, role := range ROLES {
		_, err = CreateRole(db, role["name"], role["description"])
		if err != nil {
			logger.Error("Failed to create role: " + err.Error())
		}
	}
}

func CreateRole(db *gorm.DB, name string, description string) (int, error) {
	new_role := db_tables.Role{
		Name:        name,
		Description: description,
	}
	err := db.Create(&new_role).Error
	if err != nil {
		return -1, err
	}
	return new_role.Id, nil
}
