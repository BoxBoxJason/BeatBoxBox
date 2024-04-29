package album_controller

import (
	db_model "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/utils"
	"errors"
	"mime/multipart"
	"path/filepath"
)

func PostAlbum(title string, artist_id int, illustration_file multipart.File) error {
	if title == "" {
		return errors.New("title is empty")
	}
	if AlbumExistsFromFilters(title, artist_id) {
		return errors.New("album already exists")
	}

	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err := utils.CreateNonExistingIllustrationFileName("albums")
		if err != nil {
			return err
		}
		err = utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", "albums", illustration_file_name))
		if err != nil {
			return err
		}
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	return album_model.CreateAlbum(db, title, artist_id, illustration_file_name)
}
