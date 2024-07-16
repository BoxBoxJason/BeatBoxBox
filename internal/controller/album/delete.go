package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
)

// DeleteAlbum deletes a album by its id
func DeleteAlbum(album_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return err
	}
	return album_model.DeleteAlbumFromRecord(db, &album)
}

// DeleteAlbums deletes albums by their ids
func DeleteAlbums(album_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, album_ids)
	if err != nil {
		return err
	}
	albums_pointers := make([]*db_tables.Album, len(albums))
	for i, album := range albums {
		albums_pointers[i] = &album
	}
	return album_model.DeleteAlbumsFromRecords(db, albums_pointers)
}
