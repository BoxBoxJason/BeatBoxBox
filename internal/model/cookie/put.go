package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"time"

	"gorm.io/gorm"
)

func UpdateCookieAuthToken(db *gorm.DB, cookie *db_tables.AuthCookie, new_token string) error {
	return db_model.EditRecordFields(db, cookie, map[string]interface{}{"hashed_auth_token": new_token, "expiration_date": time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix()})
}
