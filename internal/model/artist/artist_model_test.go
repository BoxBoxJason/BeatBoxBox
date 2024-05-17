package artist_model

import (
	db_model "BeatBoxBox/internal/model"
	"testing"
)

// POST FUNCTIONS TESTS
func TestArtistCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist_id, err := CreateArtist(db, "Test Artist", "default.jpg")
	if err != nil {
		t.Errorf("Error creating artist: %s", err)
	}

	if artist_id < 0 {
		t.Errorf("Expected artist Id to be a positive integer, got %d", artist_id)
	}
}

// PUT FUNCTIONS TESTS
func TestArtistUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artist_id := artist.Id
	UpdateArtist(db, artist.Id, map[string]interface{}{"name": "Updated Artist"})
	UpdateArtist(db, artist.Id, map[string]interface{}{"illustration": "update.jpg"})
	UpdateArtist(db, artist.Id, map[string]interface{}{"bio": "This is an updated test artist"})

	// Refresh the artist from the database
	artist = db_model.Artist{}
	db.Where("id = ?", artist_id).First(&artist)

	if artist.Pseudo != "Updated Artist" {
		t.Errorf("Expected artist name to be 'Updated Artist', got '%s'", artist.Pseudo)
	}
	if artist.Illustration != "update.jpg" {
		t.Errorf("Expected artist illustration to be 'update.jpg', got '%s'", artist.Illustration)
	}
	if artist.Bio != "This is an updated test artist" {
		t.Errorf("Expected artist bio to be 'This is an updated test artist', got '%s'", artist.Bio)
	}
}

// GET FUNCTIONS TESTS

func TestArtistGetFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artist, err = GetArtist(db, artist.Id)
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
	if artist.Pseudo != "Test Artist" {
		t.Errorf("Expected artist name to be 'Test Artist', got '%s'", artist.Pseudo)
	}
}

func TestArtistGetFromPseudo(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artist, err = GetArtistFromPseudo(db, artist.Pseudo)
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
	if artist.Pseudo != "Test Artist" {
		t.Errorf("Expected artist name to be 'Test Artist', got '%s'", artist.Pseudo)
	}
}

func TestArtistGetFromPartialPseudo(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artists, err := GetArtistFromPartialPseudo(db, "Test")
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
	if len(artists) < 1 {
		t.Errorf("Expected at least 1 artist, got %d", len(artists))
	}
	if artists[0].Pseudo != "Test Artist" {
		t.Errorf("Expected artist name to be 'Test Artist', got '%s'", artists[0].Pseudo)
	}
}

func TestArtistGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artists, err := GetArtistsFromFilters(db, map[string]interface{}{"pseudo": "Test Artist"})
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
	if len(artists) < 1 {
		t.Errorf("Expected at least 1 artist, got %d", len(artists))
	}
	if artists[0].Pseudo != "Test Artist" {
		t.Errorf("Expected artist name to be 'Test Artist', got '%s'", artists[0].Pseudo)
	}
}

// DELETE FUNCTIONS TESTS

func TestArtistDeleteFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artist_id := artist.Id
	err = DeleteArtist(db, artist_id)
	if err != nil {
		t.Errorf("Error deleting artist: %s", err)
	}

	artist = db_model.Artist{}
	result := db.Where("id = ?", artist_id).First(&artist)
	if result.Error == nil {
		t.Errorf("Expected artist to be deleted, got %v", artist)
	}

}

func TestArtistsDeleteFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	artist_id := artist.Id
	err = DeleteArtists(db, []int{artist_id})
	if err != nil {
		t.Errorf("Error deleting artist: %s", err)
	}

	artist = db_model.Artist{}
	result := db.Where("id = ?", artist_id).First(&artist)
	if result.Error == nil {
		t.Errorf("Expected artist to be deleted, got %v", artist)
	}
}
