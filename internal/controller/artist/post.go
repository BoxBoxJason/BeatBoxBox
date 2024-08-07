package artist_controller

import (
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

func PostArtist(pseudo string, genres []string, bio string, birthdate string, illustration_file *multipart.FileHeader) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)

	// Check if the pseudo is empty or already taken
	if pseudo == "" {
		return []byte{}, custom_errors.NewBadRequestError("artist pseudo is empty")
	}
	if IsPseudoTaken(pseudo) {
		return []byte{}, custom_errors.NewConflictError("artist already exists")
	}

	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "artists")
	if err != nil {
		return []byte{}, custom_errors.NewInternalServerError("could not upload illustration")
	}

	artist, err := artist_model.CreateArtist(db, pseudo, genres, bio, birthdate, illustration_file_name)
	if err != nil {
		return []byte{}, err
	}
	return ConvertArtistToJSON(&artist)
}
