package user_model

import (
	db_tables "BeatBoxBox/internal/model"
	role_model "BeatBoxBox/internal/model/roles"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"

	"gorm.io/gorm"
)

// PostUser creates a new user in the database
func CreateUser(db *gorm.DB, pseudo string, email string, hashed_password string, bio string, avatar_file_name string, roles_names ...string) (db_tables.User, error) {
	if len(roles_names) == 0 {
		roles_names = append(roles_names, "user")
	}
	roles, err := role_model.GetRolesByName(db, roles_names)
	if err != nil {
		return db_tables.User{}, err
	} else if len(roles) != len(roles_names) {
		return db_tables.User{}, httputils.NewNotFoundError(fmt.Sprintf("some roles not found: %v", roles_names))
	}

	new_user := db_tables.User{
		Pseudo:         pseudo,
		Email:          email,
		HashedPassword: hashed_password,
		Bio:            bio,
		Illustration:   avatar_file_name,
		Roles:          roles,
	}
	err = db.Create(&new_user).Error
	if err != nil {
		return db_tables.User{}, err
	}
	return new_user, nil
}
