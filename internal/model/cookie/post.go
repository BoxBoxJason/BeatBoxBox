package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"time"

	"gorm.io/gorm"
)

func CreateCookie(db *gorm.DB, hashed_token string, user_id int) (int, error) {
	expire_datetime := time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix()
	new_cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		UserId:          user_id,
		ExpirationDate:  expire_datetime,
	}
	err := db.Create(&new_cookie).Error
	if err != nil {
		return -1, err
	}
	return new_cookie.Id, nil
}
