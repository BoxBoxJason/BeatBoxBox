package user_model

import (
	db_model "BeatBoxBox/internal/model"
	role_model "BeatBoxBox/internal/model/roles"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"

	"gorm.io/gorm"
)

// PostUser creates a new user in the database
func CreateUser(db *gorm.DB, pseudo string, email string, hashed_password string, avatar_file_name string, roles_names ...string) (int, error) {
	if len(roles_names) == 0 {
		roles_names = append(roles_names, "user")
	}
	roles, err := role_model.GetRolesByName(db, roles_names)
	if err != nil {
		return -1, err
	} else if len(roles) != len(roles_names) {
		return -1, httputils.NewNotFoundError(fmt.Sprintf("some roles not found: %v", roles_names))
	}

	new_user := db_model.User{
		Pseudo:         pseudo,
		Email:          email,
		HashedPassword: hashed_password,
		Illustration:   avatar_file_name,
		Roles:          roles,
	}
	err = db.Create(&new_user).Error
	if err != nil {
		return -1, err
	}
	return new_user.Id, nil
}
