package album_controller

import (
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

func PostAlbum(title string, artists_ids []int, description string, illustration_file *multipart.File) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)
	if album_model.AlbumAlreadyExists(db, title, artists_ids) {
		return -1, custom_errors.NewConflictError("album already exists")
	}
	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "albums")
	if err != nil {
		return -1, custom_errors.NewInternalServerError("could not upload illustration")
	}

	return album_model.CreateAlbum(db, title, illustration_file_name)
}
