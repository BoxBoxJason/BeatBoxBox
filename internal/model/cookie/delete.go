package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	"time"

	"gorm.io/gorm"
)

// DeleteExpiredCookies deletes all the cookies that are expired
func DeleteExpiredTokens(db *gorm.DB) error {
	current_datetime := time.Now().Unix()
	return db.Where("expiration_date < ?", current_datetime).Delete(&db_tables.AuthCookie{}).Error
}

// DeleteAuthToken deletes the cookie with the given id
func DeleteAuthToken(db *gorm.DB, cookie *db_tables.AuthCookie) error {
	return db.Delete(cookie).Error
}
