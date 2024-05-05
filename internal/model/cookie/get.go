package cookie_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

func GetUserCookies(db *gorm.DB, user_id int) ([]db_model.AuthCookie, error) {
	var cookies []db_model.AuthCookie
	err := db.Where("user_id = ?", user_id).Find(&cookies).Error
	return cookies, err
}
