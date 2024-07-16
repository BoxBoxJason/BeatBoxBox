package album_controller

import (
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
)

func PostAlbum(title string, artists_ids []int, description string, illustration_file_name string) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)
	if album_model.AlbumAlreadyExists(db, title, artists_ids) {
		return -1, custom_errors.NewBadRequestError("album already exists")
	}

	return album_model.CreateAlbum(db, title, illustration_file_name)
}
