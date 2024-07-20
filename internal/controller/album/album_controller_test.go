package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"testing"
)

func TestPostAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 16",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	id, err := PostAlbum("Test Album 13", []int{artist.Id}, "description", "illustration_file_name")
	if err != nil {
		t.Error(err)
	} else if id < 0 {
		t.Error("id is negative")
	}
}

func TestDeleteAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 17",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	album := db_tables.Album{
		Title: "Test Album 14",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	album_id := album.Id
	err = DeleteAlbum(album_id)
	if err != nil {
		t.Error(err)
	}
	album = db_tables.Album{}
	err = db.First(&album, album_id).Error
	if err == nil {
		t.Error("album is not deleted")
	}
}

func TestDeleteAlbums(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 18",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	album := db_tables.Album{
		Title: "Test Album 15",
	}
	err = db.Create(&album).Error
	album_id := album.Id
	err = DeleteAlbums([]int{album_id})
	if err != nil {
		t.Error(err)
	}
	album = db_tables.Album{}
	err = db.First(&album, album_id).Error
	if err == nil {
		t.Error("album1 is not deleted")
	}
}

func TestUpdateAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 19",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	album := db_tables.Album{
		Title: "Test Album 16",
	}
	err = db.Create(&album).Error
	album_id := album.Id
	err = UpdateAlbum(album_id, map[string]interface{}{
		"description":  "Updated description",
		"illustration": "Updated illustration",
	})
	if err != nil {
		t.Error(err)
	}
	album = db_tables.Album{}
	err = db.First(&album, album_id).Error
	if err != nil {
		t.Error(err)
	} else if album.Description != "Updated description" {
		t.Error("description is not updated")
	} else if album.Illustration != "Updated illustration" {
		t.Error("illustration is not updated")
	}
}

func TestAlbumExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 19",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	if !AlbumExists(album.Id) {
		t.Error("album does not exist")
	}
}

func TestAlbumsExist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 20",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	if !AlbumsExist([]int{album.Id}) {
		t.Error("album does not exist")
	}
}

func TestGetAlbumJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 22",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	_, err = GetAlbumJSON(album.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAlbumsJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	album := db_tables.Album{
		Title: "Test Album 26",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	albums, err := GetAlbumsJSON([]int{album.Id})
	if err != nil {
		t.Error(err)
	} else if len(albums) < 1 {
		t.Error("album is not returned as expected")
	}
}

func TestGetMusicsPathFromAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 23",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	album_uint := uint(album.Id)
	music := db_tables.Music{
		Title:   "Test Music 24",
		Path:    "test_music_24.mp3",
		AlbumId: &album_uint,
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}

	_, paths, err := GetMusicsPathFromAlbum(album.Id)
	if err != nil {
		t.Error(err)
	} else if len(paths) != 1 {
		t.Error("music path is not returned as expected")
	}
}

func TestGetMusicsPathFromAlbums(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 24",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	album_uint := uint(album.Id)
	music := db_tables.Music{
		Title:   "Test Music 25",
		Path:    "test_music_25.mp3",
		AlbumId: &album_uint,
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}

	paths, err := GetMusicsPathFromAlbums([]int{album.Id})
	if err != nil {
		t.Error(err)
	} else if len(paths) != 1 {
		t.Error("music path is not returned as expected (invalid number of albums)")
	} else if len(paths["Test Album 24"]) < 1 {
		t.Error("music path is not returned as expected (invalid number of paths), got: ", paths)
	}
}

func TestGetAlbumsJSONFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	album := db_tables.Album{
		Title: "Test Album 25",
	}
	err = db.Create(&album).Error
	if err != nil {
		t.Error(err)
	}
	albums, err := GetAlbumsJSONFromPartialTitle("Test Album")
	if err != nil {
		t.Error(err)
	} else if len(albums) < 1 {
		t.Error("album is not returned as expected")
	}
}
