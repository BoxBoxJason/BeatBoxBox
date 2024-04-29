package artist_controller

import (
	db_model "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/utils"
	"errors"
	"mime/multipart"
	"path/filepath"
)

func PostArtist(pseudo string, illustration_file multipart.File) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	// Check if the pseudo is empty or already taken
	if pseudo == "" {
		return errors.New("artist pseudo is empty")
	}
	if IsPseudoTaken(pseudo) {
		return errors.New("artist already exists")
	}

	// Generate a new file name & save the illustration file if needed
	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err = utils.CreateNonExistingIllustrationFileName()
		if err != nil {
			return err
		}
		err = utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", illustration_file_name))
		if err != nil {
			return err
		}
	}
	return artist_model.CreateArtist(db, pseudo, illustration_file_name)
}
