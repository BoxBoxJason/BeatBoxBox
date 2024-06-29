package music_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"strings"
	"testing"
)

// POST FUNCTIONS TESTS
func TestMusicCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreateMusic(db, "Test Music 1", []string{"pop", "funk"}, -1, "fake.mp3", "default.jpg", -1)
	if err != nil {
		t.Errorf("Error creating music: %s", err)
	}
}

// PUT FUNCTIONS TESTS
func TestMusicUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 2",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	err = UpdateMusic(db, &music, map[string]interface{}{"title": "New Music"})
	if err != nil {
		t.Errorf("Error updating music: %s", err)
	}
	if music.Title != "New Music" {
		t.Errorf("Expected music title to be 'New Music', got '%s'", music.Title)
	}
	err = UpdateMusic(db, &music, map[string]interface{}{"path": "new_fake.mp3"})
	if err != nil {
		t.Errorf("Error updating music: %s", err)
	}
	if music.Path != "new_fake.mp3" {
		t.Errorf("Expected music path to be 'new_fake.mp3', got '%s'", music.Path)
	}
}

func TestAddArtistToMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 3",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 1",
	}
	db.Create(&artist)

	err = AddArtistsToMusic(db, &music, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error adding artist to music: %s", err)
	}

	if len(music.Artists) < 1 {
		t.Errorf("Expected at least 1 artist, got %d", len(music.Artists))
	}

}

func TestRemoveArtistFromMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 4",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	artist := db_tables.Artist{
		Pseudo: "Test Artist 2",
	}
	db.Create(&artist)

	err = AddArtistsToMusic(db, &music, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error adding artist to music: %s", err)
	}

	err = RemoveArtistsFromMusic(db, &music, []*db_tables.Artist{&artist})
	if err != nil {
		t.Errorf("Error removing artist from music: %s", err)
	}

	if len(music.Artists) > 0 {
		t.Errorf("Expected 0 artist, got %d", len(music.Artists))
	}
}

// GET FUNCTIONS TESTS

func TestGetMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 5",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music, err = GetMusic(db, music.Id)
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if music.Title != "Test Music 5" {
		t.Errorf("Expected music title to be 'Test Music 5', got '%s'", music.Title)
	}
}

func TestGetMusicsFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 6",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics := GetMusicsFromFilters(db, map[string]interface{}{"title": "Test Music 6"})
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	} else {
		if musics[0].Title != "Test Music 6" {
			t.Errorf("Expected music title to be 'Test Music', got '%s'", musics[0].Title)
		}
	}
}

func TestGetMusicsFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 7",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics, err := GetMusicsFromPartialTitle(db, "Test Music")
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	} else {
		if !strings.Contains(musics[0].Title, "Test Music") {
			t.Errorf("Expected music title to contain 'Test Music', got '%s'", musics[0].Title)
		}
	}
}

func TestGetMusics(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 8",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics, err := GetMusics(db, []int{music.Id})
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	} else {
		if !strings.Contains(musics[0].Title, "Test Music") {
			t.Errorf("Expected music title to contain 'Test Music', got '%s'", musics[0].Title)
		}
	}
}

// DELETE FUNCTIONS TESTS

func TestMusicDeleteFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 9",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music_id := music.Id
	err = DeleteMusic(db, music_id)
	if err != nil {
		t.Errorf("Error deleting music: %s", err)
	}

	music = db_tables.Music{}
	result := db.Where("id = ?", music_id).First(&music)
	if result.Error == nil {
		t.Errorf("Expected music to be deleted, got %v", music)
	}
}

func TestMusicDeleteFromRecord(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 10",
		Path:  "fake.mp3",
	}
	db.Create(&music)
	music_id := music.Id
	err = DeleteMusicFromRecord(db, &music)
	if err != nil {
		t.Errorf("Error deleting music: %s", err)
	}

	music = db_tables.Music{}
	result := db.Where("id = ?", music_id).First(&music)
	if result.Error == nil {
		t.Errorf("Expected music to be deleted, got %v", music)
	}
}

func TestMusicsDeleteFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_tables.Music{
		Title: "Test Music 11",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music_id := music.Id
	err = DeleteMusics(db, []int{music_id})
	if err != nil {
		t.Errorf("Error deleting music: %s", err)
	}

	music = db_tables.Music{}
	result := db.Where("id = ?", music_id).First(&music)
	if result.Error == nil {
		t.Errorf("Expected music to be deleted, got %v", music)
	}
}
