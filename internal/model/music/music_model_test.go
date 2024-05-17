package music_model

import (
	db_model "BeatBoxBox/internal/model"
	"testing"
)

// POST FUNCTIONS TESTS
func TestMusicCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreateMusic(db, "Test Music", []string{"pop", "funk"}, 0, "fake.mp3", "default.jpg")
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

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	UpdateMusic(db, music.Id, map[string]interface{}{"title": "New Music"})
	UpdateMusic(db, music.Id, map[string]interface{}{"path": "new_fake.mp3"})

	music_id := music.Id
	music = db_model.Music{}
	db.Where("id = ?", music_id).First(&music)

	if music.Title != "New Music" {
		t.Errorf("Expected music title to be 'New Music', got '%s'", music.Title)
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

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	err = AddArtistsToMusic(db, music.Id, []int{artist.Id})
	if err != nil {
		t.Errorf("Error adding artist to music: %s", err)
	}

	music = db_model.Music{}
	db.Where("id = ?", music.Id).Preload("Artists").First(&music)

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

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	artist := db_model.Artist{
		Pseudo: "Test Artist",
	}
	db.Create(&artist)

	err = AddArtistsToMusic(db, music.Id, []int{artist.Id})
	if err != nil {
		t.Errorf("Error adding artist to music: %s", err)
	}

	err = RemoveArtistsFromMusic(db, music.Id, []int{artist.Id})
	if err != nil {
		t.Errorf("Error removing artist from music: %s", err)
	}

	music = db_model.Music{}
	db.Where("id = ?", music.Id).Preload("Artists").First(&music)

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

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music, err = GetMusic(db, music.Id)
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if music.Title != "Test Music" {
		t.Errorf("Expected music title to be 'Test Music', got '%s'", music.Title)
	}
}

func TestGetMusicsFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics, err := GetMusicsFromFilters(db, map[string]interface{}{"title": "Test Music"})
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	}

	if musics[0].Title != "Test Music" {
		t.Errorf("Expected music title to be 'Test Music', got '%s'", musics[0].Title)
	}
}

func TestGetMusicsFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics, err := GetMusicsFromPartialTitle(db, "Test")
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	}

	if musics[0].Title != "Test Music" {
		t.Errorf("Expected music title to be 'Test Music', got '%s'", musics[0].Title)
	}
}

func TestGetMusics(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	musics, err := GetMusics(db, []int{music.Id})
	if err != nil {
		t.Errorf("Error getting music: %s", err)
	}
	if len(musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(musics))
	}

	if musics[0].Title != "Test Music" {
		t.Errorf("Expected music title to be 'Test Music', got '%s'", musics[0].Title)
	}
}

// DELETE FUNCTIONS TESTS

func TestMusicDeleteFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music_id := music.Id
	err = DeleteMusic(db, music_id)
	if err != nil {
		t.Errorf("Error deleting music: %s", err)
	}

	music = db_model.Music{}
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

	music := db_model.Music{
		Title: "Test Music",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	music_id := music.Id
	err = DeleteMusics(db, []int{music_id})
	if err != nil {
		t.Errorf("Error deleting music: %s", err)
	}

	music = db_model.Music{}
	result := db.Where("id = ?", music_id).First(&music)
	if result.Error == nil {
		t.Errorf("Expected music to be deleted, got %v", music)
	}
}
