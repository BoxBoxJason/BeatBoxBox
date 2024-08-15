package cookie_controller

import (
	cookie_model "BeatBoxBox/internal/model/cookie"
	db_model "BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"
)

func init() {
	db, err := db_model.OpenDB()
	if err == nil {
		defer db_model.CloseDB(db)
		err = cookie_model.DeleteExpiredTokens(db)
		if err != nil {
			logger.Error("Error deleting expired tokens: ", err)
		}
	} else {
		logger.Error("Error opening database: ", err)
	}
}

func DeleteMatchingAuthToken(user_id int, attempt_token string) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	auth_tokens := cookie_model.GetUserCookies(db, user_id)
	if auth_tokens == nil {
		return httputils.NewNotFoundError(fmt.Sprintf("no auth tokens found for user %d", user_id))
	}

	for _, auth_token := range auth_tokens {
		if auth_utils.CompareHash(auth_token.HashedAuthToken, attempt_token) {
			return cookie_model.DeleteAuthToken(db, &auth_token)
		}
	}

	return httputils.NewNotFoundError(fmt.Sprintf("no matching auth tokens found for user %d", user_id))
}
