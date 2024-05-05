package cookie_controller

import (
	db_model "BeatBoxBox/internal/model"
	cookie_model "BeatBoxBox/internal/model/cookie"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"errors"
)

func DeleteMatchingAuthToken(user_id int, attempt_token string) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	auth_tokens, err := cookie_model.GetUserCookies(db, user_id)
	if err != nil {
		return err
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			return cookie_model.DeleteAuthToken(db, auth_token.Id)
		}
	}

	return errors.New("no matching auth token found")
}
