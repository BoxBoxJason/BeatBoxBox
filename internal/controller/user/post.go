package user_controller

import (
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"
	"mime/multipart"
)

func PostUser(username string, email string, raw_password string, avatar_file *multipart.FileHeader) (int, error) {
	if !format_utils.CheckPseudoValidity(username) {
		return -1, httputils.NewBadRequestError("pseudo is invalid, must be between 3 and 32 characters")
	}
	if !format_utils.CheckEmailValidity(email) {
		return -1, httputils.NewBadRequestError("email is invalid, must be less than 256 characters and match the regex pattern")
	}
	if !format_utils.CheckRawPasswordValidity(raw_password) {
		return -1, httputils.NewBadRequestError("password is invalid, must be between 6 and 32 characters, contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	hashed_password, err := auth_utils.HashString(raw_password)
	if err != nil {
		return -1, err
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)
	exists, fields, err := user_model.UserAlreadyExists(db, username, email)
	if err != nil {
		return -1, err
	} else if exists {
		return -1, httputils.NewConflictError(fmt.Sprintf("user already exists with the following fields: %v", fields))
	}
	avatar_file_name, err := file_utils.UploadIllustrationToServer(avatar_file, "users")
	if err != nil {
		return -1, err
	}
	return user_model.CreateUser(db, username, email, hashed_password, avatar_file_name)
}
