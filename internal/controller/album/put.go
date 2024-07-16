package album_controller

import (
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
)

// UpdateAlbum updates an album by its id with the given map
func UpdateAlbum(album_id int, album_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return err
	}
	return album_model.UpdateAlbum(db, &album, album_map)
}
