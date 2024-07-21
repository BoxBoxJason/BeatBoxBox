package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"gorm.io/gorm"
)

func CreateCookie(db *gorm.DB, hashed_token string, user *db_tables.User) (int, error) {
	new_cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		User:            *user,
		ExpirationDate:  auth_utils.GetNewTokenExpirationTime(),
	}
	err := db.Create(&new_cookie).Error
	if err != nil {
		return -1, err
	}
	return new_cookie.Id, nil
}
