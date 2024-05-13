package artist_controller

import (
	db_model "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"errors"
	"mime/multipart"
	"path/filepath"
)

func PostArtist(pseudo string, illustration_file multipart.File) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	// Check if the pseudo is empty or already taken
	if pseudo == "" {
		return -1, errors.New("artist pseudo is empty")
	}
	if IsPseudoTaken(pseudo) {
		return -1, errors.New("artist already exists")
	}

	// Generate a new file name & save the illustration file if needed
	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err = file_utils.CreateNonExistingIllustrationFileName("artists")
		if err != nil {
			return -1, err
		}
		err = file_utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", "artists", illustration_file_name))
		if err != nil {
			return -1, err
		}
	}
	return artist_model.CreateArtist(db, pseudo, illustration_file_name)
}
