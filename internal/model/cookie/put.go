package cookie_model

import (
	db_model "BeatBoxBox/internal/model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"time"

	"gorm.io/gorm"
)

func UpdateCookieAuthToken(db *gorm.DB, cookie_id int, new_token string) error {
	return db.Model(&db_model.AuthCookie{}).Where("id = ?", cookie_id).Updates(map[string]interface{}{"hashed_auth_token": new_token, "expiration_date": time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix()}).Error
}
