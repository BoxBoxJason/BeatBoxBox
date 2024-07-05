package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	"time"

	"gorm.io/gorm"
)

func init() {
	db, err := db_model.OpenDB()
	if err == nil {
		defer db_model.CloseDB(db)
		err = DeleteExpiredTokens(db)
		if err != nil {
			logger.Error("Error deleting expired tokens: ", err)
		}
	}
}

// DeleteExpiredCookies deletes all the cookies that are expired
func DeleteExpiredTokens(db *gorm.DB) error {
	current_datetime := time.Now().Unix()
	return db.Where("expiration_date < ?", current_datetime).Delete(&db_tables.AuthCookie{}).Error
}

// DeleteAuthToken deletes the cookie with the given id
func DeleteAuthToken(db *gorm.DB, cookie *db_tables.AuthCookie) error {
	return db.Delete(cookie).Error
}
