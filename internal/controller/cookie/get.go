package cookie_controller

import (
	db_model "BeatBoxBox/internal/model"
	cookie_model "BeatBoxBox/internal/model/cookie"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"errors"
)

func CheckAuthTokenMatches(user_id int, attempt_token string) (bool, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return false, "", err
	}
	defer db_model.CloseDB(db)

	auth_tokens, err := cookie_model.GetUserCookies(db, user_id)
	if err != nil {
		return false, "", err
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			new_token, err := updateAuthTokenIfNearExpiry(auth_token)
			if err != nil {
				return false, "", err
			}
			return true, new_token, nil
		}
	}

	return false, "", nil
}

func GetMatchingAuthTokenId(user_id int, attempt_token string) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	auth_tokens, err := cookie_model.GetUserCookies(db, user_id)
	if err != nil {
		return -1, err
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			return auth_token.Id, nil
		}
	}

	return -1, errors.New("no matching auth token found in database")
}
