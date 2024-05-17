package album_model

import (
	db_model "BeatBoxBox/internal/model"
	"testing"
)

// POST FUNCTIONS TESTS
func TestAlbumCreation(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	id, err := CreateAlbum(db, "Test Album", "This is a test album")

	if err != nil {
		t.Errorf("Error creating album: %s", err)
	}
	if id < 0 {
		t.Errorf("Expected album Id to be a positive integer, got %d", id)
	}
}

// PUT FUNCTIONS TESTS
func TestAlbumUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	UpdateAlbum(db, album.Id, map[string]interface{}{"title": "Updated Album"})
	UpdateAlbum(db, album.Id, map[string]interface{}{"description": "This is an updated test album"})
	UpdateAlbum(db, album.Id, map[string]interface{}{"illustration": "update.jpg"})

	// Refresh the album from the database
	album = db_model.Album{}
	db.Where("id = ?", album.Id).First(&album)

	if album.Title != "Updated Album" {
		t.Errorf("Expected album title to be 'Updated Album', got '%s'", album.Title)
	}
	if album.Description != "This is an updated test album" {
		t.Errorf("Expected album description to be 'This is an updated test album', got '%s'", album.Description)
	}
	if album.Illustration != "update.jpg" {
		t.Errorf("Expected album illustration to be 'update.jpg', got '%s'", album.Illustration)
	}
}

func TestAddMusicsToAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)
	album_id := album.Id

	music := db_model.Music{
		Title: "Test Music",
	}
	db.Create(&music)

	err = AddMusicsToAlbum(db, album.Id, []int{music.Id})
	if err != nil {
		t.Errorf("Error adding music to album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	db.Preload("Musics").Where("id = ?", album_id).First(&album)

	if len(album.Musics) < 1 {
		t.Errorf("Expected at least 1 music in album, got %d", len(album.Musics))
	}
}

func TestRemoveMusicsFromAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)
	album_id := album.Id

	music := db_model.Music{
		Title: "Test Music",
	}
	db.Create(&music)

	err = AddMusicsToAlbum(db, album.Id, []int{music.Id})
	if err != nil {
		t.Errorf("Error adding music to album: %s", err)
	}

	err = RemoveMusicsFromAlbum(db, album.Id, []int{music.Id})
	if err != nil {
		t.Errorf("Error removing music from album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	db.Preload("Musics").Where("id = ?", album_id).First(&album)

	if len(album.Musics) > 0 {
		t.Errorf("Expected no music in album, got %d", len(album.Musics))
	}
}

func TestAddArtistToAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)
	album_id := album.Id

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	err = AddArtistToAlbum(db, album.Id, artist.Id)
	if err != nil {
		t.Errorf("Error adding artist to album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	db.Preload("Artists").Where("id = ?", album_id).First(&album)

	if len(album.Artists) < 1 {
		t.Errorf("Expected at least 1 artist in album, got %d", len(album.Artists))
	}
}

func TestRemoveArtistFromAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)
	album_id := album.Id

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	err = AddArtistToAlbum(db, album.Id, artist.Id)
	if err != nil {
		t.Errorf("Error adding artist to album: %s", err)
	}

	err = RemoveArtistFromAlbum(db, album.Id, artist.Id)
	if err != nil {
		t.Errorf("Error removing artist from album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	db.Preload("Artists").Where("id = ?", album_id).First(&album)

	if len(album.Artists) > 0 {
		t.Errorf("Expected no artist in album, got %d", len(album.Artists))
	}
}

// GET FUNCTIONS TESTS

func TestAlbumGetFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	retrieved_album, err := GetAlbum(db, album.Id)
	if err != nil {
		t.Errorf("Error retrieving album: %s", err)
	}

	if retrieved_album.Title != album.Title {
		t.Errorf("Expected album title to be '%s', got '%s'", album.Title, retrieved_album.Title)
	}
}

func TestAlbumGetFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	albums, err := GetAlbumsFromPartialTitle(db, "Test")
	if err != nil {
		t.Errorf("Error retrieving album: %s", err)
	}

	if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

func TestAlbumGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	albums, err := GetAlbumsFromFilters(db, map[string]interface{}{"title": "Test Album"})
	if err != nil {
		t.Errorf("Error retrieving album: %s", err)
	}

	if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

func TestAlbumGetFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	albums, err := GetAlbums(db, []int{album.Id})
	if err != nil {
		t.Errorf("Error retrieving album: %s", err)
	}

	if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

// DELETE FUNCTIONS TESTS

func TestAlbumDeleteFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	err = DeleteAlbum(db, album.Id)
	if err != nil {
		t.Errorf("Error deleting album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	result := db.Where("id = ?", album.Id).First(&album)
	if result.Error == nil {
		t.Errorf("Expected error when retrieving deleted album, got nil")
	}
}

func TestAlbumDeleteFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_model.Album{
		Title: "Test Album",
	}
	db.Create(&album)

	err = DeleteAlbums(db, []int{album.Id})
	if err != nil {
		t.Errorf("Error deleting album: %s", err)
	}

	// Refresh the album from the database
	album = db_model.Album{}
	result := db.Where("id = ?", album.Id).First(&album)
	if result.Error == nil {
		t.Errorf("Expected error when retrieving deleted album, got nil")
	}
}
