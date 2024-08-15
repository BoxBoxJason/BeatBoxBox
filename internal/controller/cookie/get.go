package cookie_controller

import (
	cookie_model "BeatBoxBox/internal/model/cookie"
	db_model "BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"
)

func CheckAuthTokenMatches(user_id int, attempt_token string) (bool, string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return false, "", err
	}
	defer db_model.CloseDB(db)

	auth_tokens := cookie_model.GetUserCookies(db, user_id)
	if auth_tokens == nil || len(auth_tokens) == 0 {
		return false, "", httputils.NewNotFoundError(fmt.Sprintf("no auth tokens found for user %d", user_id))
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			new_token, err := updateAuthTokenIfNearExpiry(&auth_token)
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

	auth_tokens := cookie_model.GetUserCookies(db, user_id)
	if auth_tokens == nil || len(auth_tokens) == 0 {
		return -1, httputils.NewNotFoundError(fmt.Sprintf("no auth tokens found for user %d", user_id))
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			return auth_token.Id, nil
		}
	}

	return -1, httputils.NewNotFoundError(fmt.Sprintf("no matching auth token found for user %d", user_id))
}
