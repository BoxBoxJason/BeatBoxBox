package playlist_model

import (
	db_model "BeatBoxBox/internal/model"
	"testing"
)

// POST FUNCTIONS TESTS

func TestPlaylistCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreatePlaylist(db, "Test Playlist", 0, "", "default.jpg")
	if err != nil {
		t.Errorf("Error creating playlist: %s", err)
	}
}

// PUT FUNCTIONS TESTS

func TestPlaylistUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_model.Playlist{
		Title: "Test Playlist",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	UpdatePlaylist(db, playlist_id, map[string]interface{}{"title": "New Test Playlist"})
	playlist = db_model.Playlist{}
	db.Where("id = ?", playlist_id).First(&playlist)

	if playlist.Title != "New Test Playlist" {
		t.Errorf("Expected playlist name to be 'New Test Playlist', got '%s'", playlist.Title)
	}
}

func TestAddMusicsToPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_model.Playlist{
		Title: "Test Playlist",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	music1 := db_model.Music{
		Title: "Test Music 1",
	}
	music2 := db_model.Music{
		Title: "Test Music 2",
	}
	db.Create(&music1)
	db.Create(&music2)
	music_ids := []int{music1.Id, music2.Id}

	err = AddMusicsToPlaylist(db, playlist_id, music_ids)
	if err != nil {
		t.Errorf("Error adding musics to playlist: %s", err)
	}

	playlist = db_model.Playlist{}
	db.Preload("Musics").Where("id = ?", playlist_id).First(&playlist)

	if len(playlist.Musics) < 2 {
		t.Errorf("Expected 2 musics, got %d", len(playlist.Musics))
	}
}

func TestRemoveMusicsFromPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_model.Playlist{
		Title: "Test Playlist",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	music1 := db_model.Music{
		Title: "Test Music 1",
	}
	music2 := db_model.Music{
		Title: "Test Music 2",
	}
	db.Create(&music1)
	db.Create(&music2)
	music_ids := []int{music1.Id, music2.Id}

	err = AddMusicsToPlaylist(db, playlist_id, music_ids)
	if err != nil {
		t.Errorf("Error adding musics to playlist: %s", err)
	}

	err = RemoveMusicsFromPlaylist(db, playlist_id, music_ids)
	if err != nil {
		t.Errorf("Error removing musics from playlist: %s", err)
	}

	playlist = db_model.Playlist{}
	db.Preload("Musics").Where("id = ?", playlist_id).First(&playlist)

	if len(playlist.Musics) > 0 {
		t.Errorf("Expected 0 musics, got %d", len(playlist.Musics))
	}
}

// GET FUNCTIONS TESTS

func TestGetPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_model.Playlist{
		Title: "Test Playlist",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	playlist, err = GetPlaylist(db, playlist_id)
	if err != nil {
		t.Errorf("Error getting playlist: %s", err)
	}

	if playlist.Title != "Test Playlist" {
		t.Errorf("Expected playlist name to be 'Test Playlist', got '%s'", playlist.Title)
	}
}

func TestGetPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_model.Playlist{
		Title: "Test Playlist 1",
	}
	playlist2 := db_model.Playlist{
		Title: "Test Playlist 2",
	}
	db.Create(&playlist1)
	db.Create(&playlist2)
	playlist_ids := []int{playlist1.Id, playlist2.Id}

	playlists, err := GetPlaylists(db, playlist_ids)
	if err != nil {
		t.Errorf("Error getting playlists: %s", err)
	}

	if len(playlists) != 2 {
		t.Errorf("Expected 2 playlists, got %d", len(playlists))
	}
}

func TestGetPlaylistsFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_model.Playlist{
		Title: "Test Playlist 1",
	}
	playlist2 := db_model.Playlist{
		Title: "Test Playlist 2",
	}
	db.Create(&playlist1)
	db.Create(&playlist2)

	playlists, err := GetPlaylistsFromFilters(db, map[string]interface{}{"title": "Test Playlist 1"})
	if err != nil {
		t.Errorf("Error getting playlists: %s", err)
	}

	if len(playlists) == 1 {
		t.Errorf("Expected 1 playlist, got %d", len(playlists))
	}
}

func TestGetPlaylistsFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_model.Playlist{
		Title: "Test Playlist 1",
	}
	playlist2 := db_model.Playlist{
		Title: "Test Playlist 2",
	}
	db.Create(&playlist1)
	db.Create(&playlist2)

	playlists, err := GetPlaylistsFromPartialTitle(db, "Test")
	if err != nil {
		t.Errorf("Error getting playlists: %s", err)
	}

	if len(playlists) < 2 {
		t.Errorf("Expected at least 2 playlists, got %d", len(playlists))
	}
}

// DELETE FUNCTIONS TESTS

func TestPlaylistDelete(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_model.Playlist{
		Title: "Test Playlist",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	DeletePlaylist(db, playlist_id)
	playlist = db_model.Playlist{}
	result := db.Where("id = ?", playlist_id).First(&playlist)

	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestPlaylistsDeleteFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_model.Playlist{
		Title: "Test Playlist 1",
	}
	playlist2 := db_model.Playlist{
		Title: "Test Playlist 2",
	}
	db.Create(&playlist1)
	db.Create(&playlist2)
	playlist_ids := []int{playlist1.Id, playlist2.Id}

	DeletePlaylists(db, playlist_ids)
	playlists := []db_model.Playlist{}
	result := db.Where("id IN ?", playlist_ids).Find(&playlists)

	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}
