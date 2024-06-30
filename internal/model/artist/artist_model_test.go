package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS TESTS
func TestArtistCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist_id, err := CreateArtist(db, "Test Artist 5", "default.jpg")
	if err != nil {
		t.Errorf("Error creating artist: %s", err)
	} else if artist_id < 0 {
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

	artist := db_tables.Artist{
		Pseudo: "Test Artist 6",
	}
	db.Create(&artist)

	err = UpdateArtist(db, &artist, map[string]interface{}{"pseudo": "Updated Artist 6", "illustration": "update.jpg", "bio": "This is an updated test artist"})
	if err != nil {
		t.Errorf("Error updating artist: %s", err)
	}

	if artist.Pseudo != "Updated Artist 6" {
		t.Errorf("Expected artist name to be 'Updated Artist', got '%s'", artist.Pseudo)
	}
	if artist.Illustration != "update.jpg" {
		t.Errorf("Expected artist illustration to be 'update.jpg', got '%s'", artist.Illustration)
	}
	if artist.Bio != "This is an updated test artist" {
		t.Errorf("Expected artist bio to be 'This is an updated test artist', got '%s'", artist.Bio)
	}
}

func TestArtistAddMusicsToArtist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 7",
	}
	db.Create(&artist)

	music := db_tables.Music{
		Title: "Test Music 17",
		Path:  "test.mp3",
	}
	db.Create(&music)

	err = AddMusicsToArtist(db, &artist, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error adding music to artist: %s", err)
	} else if len(artist.Musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(artist.Musics))
	}
}

func TestArtistRemoveMusicsFromArtist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 8",
	}
	db.Create(&artist)

	music := db_tables.Music{
		Title: "Test Music 18",
		Path:  "test.mp3",
	}
	db.Create(&music)
	artist.Musics = append(artist.Musics, music)

	err = RemoveMusicsFromArtist(db, &artist, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error removing music from artist: %s", err)
	}
	if len(artist.Musics) > 0 {
		t.Errorf("Expected 0 music, got %d", len(artist.Musics))
	}
}

// GET FUNCTIONS TESTS

func TestArtistGetFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 9",
	}
	db.Create(&artist)

	artist, err = GetArtist(db, artist.Id)
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
}

func TestArtistGetFromPartialPseudo(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 10",
	}
	db.Create(&artist)

	artists, err := GetArtistsFromPartialPseudo(db, map[string]interface{}{}, "Test Artist")
	if err != nil {
		t.Errorf("Error getting artist: %s", err)
	}
	if len(artists) < 1 {
		t.Errorf("Expected at least 1 artist, got %d", len(artists))
	}
}

func TestArtistGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 11",
	}
	db.Create(&artist)

	artists := GetArtistsFromFilters(db, map[string]interface{}{"pseudo": "Test Artist 11"})
	if artists == nil {
		t.Errorf("Expected at least 1 artist, got 0")
	}
	if len(artists) < 1 {
		t.Errorf("Expected at least 1 artist, got %d", len(artists))
	}
}

// DELETE FUNCTIONS TESTS

func TestArtistDeleteFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 12",
	}
	db.Create(&artist)

	artist_id := artist.Id
	err = DeleteArtistFromRecord(db, artist)
	if err != nil {
		t.Errorf("Error deleting artist: %s", err)
	}

	artist = db_tables.Artist{}
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

	artist := db_tables.Artist{
		Pseudo: "Test Artist 13",
	}
	db.Create(&artist)

	artist_id := artist.Id
	err = DeleteArtistsFromRecords(db, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error deleting artist: %s", err)
	}

	artist = db_tables.Artist{}
	result := db.Where("id = ?", artist_id).First(&artist)
	if result.Error == nil {
		t.Errorf("Expected artist to be deleted, got %v", artist)
	}
}
