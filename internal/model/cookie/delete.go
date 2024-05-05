package cookie_model

import (
	db_model "BeatBoxBox/internal/model"
	"time"

	"gorm.io/gorm"
)

func init() {
	db, err := db_model.OpenDB()
	if err == nil {
		defer db_model.CloseDB(db)
		DeleteExpiredTokens(db)
	}

}

// DeleteExpiredCookies deletes all the cookies that are expired
func DeleteExpiredTokens(db *gorm.DB) error {
	current_datetime := time.Now().Unix()
	return db.Where("expiration_date < ?", current_datetime).Delete(&db_model.AuthCookie{}).Error
}

// DeleteAuthToken deletes the cookie with the given id
func DeleteAuthToken(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&db_model.AuthCookie{}).Error
}
