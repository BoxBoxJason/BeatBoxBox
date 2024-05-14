package album_controller

import (
	db_model "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
)

func UpdateAlbum(album_id int, album_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return album_model.UpdateAlbum(db, album_id, album_map)
}

func AddMusicsToAlbum(album_id int, musics_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return album_model.AddMusicsToAlbum(db, album_id, musics_ids)
}

func RemoveMusicsFromAlbum(album_id int, musics_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return album_model.RemoveMusicsFromAlbum(db, album_id, musics_ids)
}
