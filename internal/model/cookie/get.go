package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

func GetUserCookies(db *gorm.DB, user_id int) []db_tables.AuthCookie {
	cookies := db_model.GetRecordsByField(db, &db_tables.AuthCookie{}, "user_id", user_id)
	if cookies == nil {
		return nil
	}
	user_cookies := make([]db_tables.AuthCookie, len(cookies))
	for i, cookie := range cookies {
		user_cookies[i] = cookie.(db_tables.AuthCookie)
	}

	return user_cookies
}
