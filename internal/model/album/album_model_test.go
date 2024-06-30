package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS TESTS
func TestAlbumCreation(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	id, err := CreateAlbum(db, "Test Album 1", "fake.jpeg")

	if err != nil {
		t.Errorf("Error creating album: %s", err)
	} else if id < 0 {
		t.Errorf("Error creating album: id is negative")
	}
}

// PUT FUNCTIONS TESTS
func TestAlbumUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 2",
	}
	db.Create(&album)

	err = UpdateAlbum(db, &album, map[string]interface{}{"title": "Updated Album 2", "description": "This is an updated test album", "illustration": "update.jpg"})
	if err != nil {
		t.Errorf("Error updating album: %s", err)
	}

	if album.Title != "Updated Album 2" {
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

	album := db_tables.Album{
		Title: "Test Album 3",
	}
	db.Create(&album)

	music := db_tables.Music{
		Title: "Test Music 15",
		Path:  "test.mp3",
	}
	db.Create(&music)

	err = AddMusicsToAlbum(db, &album, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error adding music to album: %s", err)
	}

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

	album := db_tables.Album{
		Title: "Test Album 4",
	}
	db.Create(&album)

	music := db_tables.Music{
		Title: "Test Music 16",
		Path:  "test.mp3",
	}
	db.Create(&music)

	album.Musics = append(album.Musics, music)

	err = RemoveMusicsFromAlbum(db, &album, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error removing music from album: %s", err)
	}

	if len(album.Musics) > 0 {
		t.Errorf("Expected no music in album, got %d", len(album.Musics))
	}
}

func TestAddArtistsToAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 5",
	}
	db.Create(&album)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 3",
	}
	db.Create(&artist)

	err = AddArtistsToAlbum(db, &album, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error adding artist to album: %s", err)
	}

	if len(album.Artists) < 1 {
		t.Errorf("Expected at least 1 artist in album, got %d", len(album.Artists))
	}
}

func TestRemoveArtistsFromAlbum(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 6",
	}
	db.Create(&album)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 4",
	}
	db.Create(&artist)

	album.Artists = append(album.Artists, artist)

	err = RemoveArtistsFromAlbum(db, &album, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error removing artist from album: %s", err)
	}

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

	album := db_tables.Album{
		Title: "Test Album 7",
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

	album := db_tables.Album{
		Title: "Test Album 8",
	}
	db.Create(&album)

	albums := GetAlbumsFromPartialTitle(db, map[string]interface{}{}, "Test Album")
	if albums == nil {
		t.Errorf("Error retrieving album")
	} else if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

func TestAlbumGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 9",
	}
	db.Create(&album)

	albums := GetAlbumsFromFilters(db, map[string]interface{}{"title": "Test Album 9"})
	if albums == nil {
		t.Errorf("Error retrieving album")
	} else if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

func TestAlbumGetFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 10",
	}
	db.Create(&album)

	albums, err := GetAlbums(db, []int{album.Id})
	if err != nil {
		t.Errorf("Error retrieving album: %s", err)
	} else if len(albums) < 1 {
		t.Errorf("Expected at least 1 album, got %d", len(albums))
	}
}

// DELETE FUNCTIONS TESTS

func TestAlbumDeleteFromRecord(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 11",
	}
	db.Create(&album)
	album_id := album.Id
	err = DeleteAlbumFromRecord(db, album)
	if err != nil {
		t.Errorf("Error deleting album: %s", err)
	}

	// Refresh the album from the database
	album = db_tables.Album{}
	result := db.Where("id = ?", album_id).First(&album)
	if result.Error == nil {
		t.Errorf("Expected error when retrieving deleted album, got nil")
	}
}

func TestAlbumDeleteFromRecords(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	album := db_tables.Album{
		Title: "Test Album 12",
	}
	db.Create(&album)
	album_id := album.Id
	err = DeleteAlbumsFromRecords(db, []db_tables.Album{album})
	if err != nil {
		t.Errorf("Error deleting album: %s", err)
	}

	// Refresh the album from the database
	album = db_tables.Album{}
	result := db.Where("id = ?", album_id).First(&album)
	if result.Error == nil {
		t.Errorf("Expected error when retrieving deleted album, got nil")
	}
}
