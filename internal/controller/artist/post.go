package artist_controller

import (
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

func PostArtist(pseudo string, bio string, illustration_file *multipart.File) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	// Check if the pseudo is empty or already taken
	if pseudo == "" {
		return -1, custom_errors.NewBadRequestError("artist pseudo is empty")
	}
	if IsPseudoTaken(pseudo) {
		return -1, custom_errors.NewConflictError("artist already exists")
	}

	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "artists")
	if err != nil {
		return -1, custom_errors.NewInternalServerError("could not upload illustration")
	}

	return artist_model.CreateArtist(db, pseudo, bio, illustration_file_name)
}
