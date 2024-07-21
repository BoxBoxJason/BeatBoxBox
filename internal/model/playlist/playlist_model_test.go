package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS TESTS

func TestPlaylistCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreatePlaylist(db, "Test Playlist 3", []*db_tables.User{}, "", "default.jpg")
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

	playlist := db_tables.Playlist{
		Title: "Test Playlist 4",
	}
	db.Create(&playlist)

	err = UpdatePlaylist(db, &playlist, map[string]interface{}{"title": "New Test Playlist 4"})
	if err != nil {
		t.Errorf("Error updating playlist: %s", err)
	}

	if playlist.Title != "New Test Playlist 4" {
		t.Errorf("Expected playlist name to be 'New Test Playlist 4', got '%s'", playlist.Title)
	}
}

func TestAddMusicsToPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_tables.Playlist{
		Title: "Test Playlist 5",
	}
	db.Create(&playlist)

	music1 := db_tables.Music{
		Title: "Test Music 19",
		Path:  "test.mp3",
	}

	db.Create(&music1)

	err = AddMusicsToPlaylist(db, &playlist, []*db_tables.Music{&music1})
	if err != nil {
		t.Errorf("Error adding musics to playlist: %s", err)
	} else if len(playlist.Musics) < 1 {
		t.Errorf("Expected at least 1 music, got %d", len(playlist.Musics))
	}
}

func TestRemoveMusicsFromPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_tables.Playlist{
		Title: "Test Playlist 6",
	}
	db.Create(&playlist)

	music1 := db_tables.Music{
		Title: "Test Music 20",
		Path:  "test.mp3",
	}
	db.Create(&music1)

	playlist.Musics = append(playlist.Musics, music1)

	err = RemoveMusicsFromPlaylist(db, &playlist, []*db_tables.Music{&music1})
	if err != nil {
		t.Errorf("Error removing musics from playlist: %s", err)
	} else if len(playlist.Musics) > 0 {
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

	playlist := db_tables.Playlist{
		Title: "Test Playlist 7",
	}
	db.Create(&playlist)

	playlist, err = GetPlaylist(db, playlist.Id)
	if err != nil {
		t.Errorf("Error getting playlist: %s", err)
	}
}

func TestGetPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_tables.Playlist{
		Title: "Test Playlist 8",
	}

	db.Create(&playlist1)

	playlists, err := GetPlaylists(db, []int{playlist1.Id})
	if err != nil {
		t.Errorf("Error getting playlists: %s", err)
	}
	if len(playlists) == 0 {
		t.Errorf("Expected at least 1 playlist, got 0")
	}
}

func TestGetPlaylistsFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_tables.Playlist{
		Title: "Test Playlist 9",
	}
	db.Create(&playlist1)

	playlists := GetPlaylistsFromFilters(db, map[string]interface{}{"title": "Test Playlist 9"})
	if len(playlists) != 1 {
		t.Errorf("Expected 1 playlist, got %d", len(playlists))
	}
}

func TestGetPlaylistsFromPartialTitle(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_tables.Playlist{
		Title: "Test Playlist 10",
	}
	db.Create(&playlist1)

	playlists := GetPlaylistsFromPartialTitle(db, map[string]interface{}{}, "Test")
	if playlists == nil {
		t.Errorf("Expected a playlist, got nil")
	} else if len(playlists) == 0 {
		t.Errorf("Expected at least 1 playlist, got 0")
	}
}

// DELETE FUNCTIONS TESTS

func TestPlaylistDelete(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist := db_tables.Playlist{
		Title: "Test Playlist 11",
	}
	db.Create(&playlist)
	playlist_id := playlist.Id

	err = DeletePlaylist(db, &playlist)
	playlist = db_tables.Playlist{}
	result := db.Where("id = ?", playlist_id).First(&playlist)

	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestPlaylistsDelete(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	playlist1 := db_tables.Playlist{
		Title: "Test Playlist 12",
	}
	db.Create(&playlist1)
	playlist1_id := playlist1.Id
	err = DeletePlaylists(db, []*db_tables.Playlist{&playlist1})
	if err != nil {
		t.Errorf("Error deleting playlists: %s", err)
	}

	playlist1 = db_tables.Playlist{}
	result := db.Where("id = ?", playlist1_id).First(&playlist1)
	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestPlaylistAlreadyExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 47",
		Email:           "Test Email 47",
		Hashed_password: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title:  "Test Playlist 30",
		Owners: []db_tables.User{user},
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}

	if !PlaylistAlreadyExists(db, "Test Playlist 30", []int{user.Id}) {
		t.Errorf("Expected playlist to exist, got false")
	}
}
