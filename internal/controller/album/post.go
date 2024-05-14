package album_controller

import (
	db_model "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	custom_errors "BeatBoxBox/pkg/errors"
)

func PostAlbum(title string, artists_ids []int, description string, illustration_file_name string) (int, error) {
	if title == "" {
		return -1, custom_errors.NewBadRequestError("title is empty")
	}
	if AlbumExistsFromFilters(title, artists_ids) {
		return -1, custom_errors.NewBadRequestError("album already exists")
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return album_model.CreateAlbum(db, title, illustration_file_name)
}
