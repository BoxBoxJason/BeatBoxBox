package cookie_model

import (
	db_model "BeatBoxBox/internal/model"
	"time"

	"gorm.io/gorm"
)

const DEFAULT_TOKEN_EXPIRATION = 48 * time.Hour

func CreateCookie(db *gorm.DB, hashed_token string, user_id int) (int, error) {
	expire_datetime := time.Now().Add(DEFAULT_TOKEN_EXPIRATION).Unix()
	new_cookie := db_model.AuthCookie{
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
