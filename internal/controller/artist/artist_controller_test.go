package artist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS
func TestPostArtist(t *testing.T) {
	id, err := PostArtist("Test Artist 23", "description", nil)
	if err != nil {
		t.Error(err)
	} else if id < 0 {
		t.Error("id is negative")
	}
}

// DELETE FUNCTIONS
func TestDeleteArtist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 24",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	err = DeleteArtist(artist_id)
	if err != nil {
		t.Error(err)
	}

	if db.First(&db_tables.Artist{}, artist_id).Error == nil {
		t.Error("artist is not deleted")
	}
}

func TestDeleteArtists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 25",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	err = DeleteArtists([]int{artist_id})
	if err != nil {
		t.Error(err)
	}

	if db.First(&db_tables.Artist{}, artist_id).Error == nil {
		t.Error("artist is not deleted")
	}
}

// PUT FUNCTIONS
func TestUpdateArtist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 26",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	err = UpdateArtist(artist_id, map[string]interface{}{"pseudo": "Test Artist 26 New"})
	if err != nil {
		t.Error(err)
	}

	artist = db_tables.Artist{}
	err = db.First(&artist, artist_id).Error
	if err != nil || artist.Pseudo != "Test Artist 26 New" {
		t.Error("artist is not updated")
	}
}

// GET FUNCTIONS
func TestGetArtistJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 27",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	_, err = GetArtistJSON(artist_id)
	if err != nil {
		t.Error(err)
	}
}

func TestGetArtistsJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 28",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	_, err = GetArtistsJSON([]int{artist_id})
	if err != nil {
		t.Error(err)
	}
}

func TestGetArtistsJSONFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 32",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	_, err = GetArtistsJSONFromFilters(map[string]interface{}{"pseudo": "Test Artist 32"})
	if err != nil {
		t.Error(err)
	}
}

func TestArtistExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 29",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	if !ArtistExists(artist_id) {
		t.Error("artist does not exist")
	}
}

func TestArtistsExist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 30",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	artist_id := artist.Id
	if !ArtistsExist([]int{artist_id}) {
		t.Error("At least one artist does not exist")
	}
}

func TestIsPseudoTaken(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 31",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	if IsPseudoTaken("Test Artist 31") {
		t.Error("pseudo is not taken")
	}
}
