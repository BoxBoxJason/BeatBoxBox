package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS
func TestPostPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 29",
		Hashed_password: "hashed_password",
		Email:           "Test Email 29",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	_, err = PostPlaylist("Test Playlist 13", user.Id, "description", nil)
	if err != nil {
		t.Error(err)
	}
}

// DELETE FUNCTIONS
func TestDeletePlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 14",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	playlist_id := playlist.Id
	err = DeletePlaylist(playlist_id)
	if err != nil {
		t.Error(err)
	}
	err = db.First(&playlist, playlist_id).Error
	if err == nil {
		t.Error("Playlist not deleted")
	}
}

func TestDeletePlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 15",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	playlist_id := playlist.Id
	err = DeletePlaylists([]int{playlist_id})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&playlist, playlist_id).Error
	if err == nil {
		t.Error("Playlist not deleted")
	}
}

// GET FUNCTIONS
func TestPlaylistExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 16",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	if !PlaylistExists(playlist.Id) {
		t.Error("Playlist not found")
	}
}

func TestPlaylistsExist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 17",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	if !PlaylistsExist([]int{playlist.Id}) {
		t.Error("Playlist not found")
	}
}

func TestPlaylistAlreadyExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 28",
		Hashed_password: "hashed_password",
		Email:           "Test Email 28",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title:   "Test Playlist 18",
		Creator: user,
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	if !PlaylistAlreadyExists(playlist.Title, user.Id) {
		t.Error("Playlist not found")
	}
}

func TestGetPlaylistJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 19",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	_, err = GetPlaylistJSON(playlist.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPlaylistsJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 20",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	_, err = GetPlaylistsJSON([]int{playlist.Id})
	if err != nil {
		t.Error(err)
	}
}

func TestGetMusicsPathFromPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 38",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title:  "Test Playlist 21",
		Musics: []db_tables.Music{music},
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	_, musics_path, err := GetMusicsPathFromPlaylist(playlist.Id)
	if err != nil {
		t.Error(err)
	} else if len(musics_path) == 0 {
		t.Error("No musics path found")
	}
}

func TestGetMusicsPathFromPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 39",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title:  "Test Playlist 22",
		Musics: []db_tables.Music{music},
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	musics_map, err := GetMusicsPathFromPlaylists([]int{playlist.Id})
	if err != nil {
		t.Error(err)
	} else if len(musics_map) == 0 {
		t.Error("No playlist found")
	} else if len(musics_map[playlist.Title]) == 0 {
		t.Error("No musics path found")
	}
}

// PUT FUNCTIONS
func TestUpdatePlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 23",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	playlist_map := map[string]interface{}{
		"title": "Test Playlist 23 updated",
	}
	err = UpdatePlaylist(playlist.Id, playlist_map)
	if err != nil {
		t.Error(err)
	}
	err = db.First(&playlist, playlist.Id).Error
	if err != nil {
		t.Error(err)
	} else if playlist.Title != "Test Playlist 23 updated" {
		t.Error("Playlist not updated")
	}
}

func TestAddMusicsToPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 40",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title: "Test Playlist 24",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	err = AddMusicsToPlaylist(playlist.Id, []int{music.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Musics").Where("id = ?", playlist.Id).First(&playlist).Error
	if err != nil {
		t.Error(err)
	} else if len(playlist.Musics) == 0 {
		t.Error("Music not added to playlist")
	}
}

func TestRemoveMusicsFromPlaylist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 41",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title:  "Test Playlist 25",
		Musics: []db_tables.Music{music},
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	err = RemoveMusicsFromPlaylist(playlist.Id, []int{music.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Musics").Where("id = ?", playlist.Id).First(&playlist).Error
	if err != nil {
		t.Error(err)
	} else if len(playlist.Musics) != 0 {
		t.Error("Music not removed from playlist")
	}
}
