package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"strings"
	"testing"
)

// POST FUNCTIONS
func TestPostMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	artist := db_tables.Artist{
		Pseudo: "Test Artist 34",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	_, err = PostMusic("Test Music 26", []string{"pop", "funk"}, "Test Lyrics 26", -1, nil, nil, []int{artist.Id})
	if err != nil {
		t.Error(err)
	}
}

// DELETE FUNCTIONS
func TestDeleteMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 27",
		Path:  "test_music_27.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_id := music.Id
	err = DeleteMusic(music_id)
	if err != nil {
		t.Error(err)
	}
	err = db.First(&music, music_id).Error
	if err == nil {
		t.Error("Music was not deleted")
	}
}

func TestDeleteMusics(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 28",
		Path:  "test_music_28.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_id := music.Id
	err = DeleteMusics([]int{music_id})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&music, music_id).Error
	if err == nil {
		t.Error("Music was not deleted")
	}
}

// PUT FUNCTIONS
func TestUpdateMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 29",
		Path:  "test_music_29.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_id := music.Id
	err = UpdateMusic(music_id, map[string]interface{}{"title": "Test Music 29 Updated"})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&music, music_id).Error
	if err != nil {
		t.Error(err)
	} else if music.Title != "Test Music 29 Updated" {
		t.Error("Music was not updated")
	}
}

func TestAddArtistsToMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 30",
		Path:  "test_music_30.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_id := music.Id
	artist := db_tables.Artist{
		Pseudo: "Test Artist 35",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	err = AddArtistsToMusic(music_id, []int{artist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Artists").First(&music, music_id).Error
	if err != nil {
		t.Error(err)
	} else if len(music.Artists) != 1 {
		t.Error("Artist was not added to music")
	}
}

func TestRemoveArtistsFromMusic(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	artist := db_tables.Artist{
		Pseudo: "Test Artist 36",
	}
	err = db.Create(&artist).Error
	if err != nil {
		t.Error(err)
	}
	music := db_tables.Music{
		Title:   "Test Music 31",
		Path:    "test_music_31.mp3",
		Artists: []db_tables.Artist{artist},
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_id := music.Id
	err = RemoveArtistsFromMusic(music_id, []int{artist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Artists").First(&music, music_id).Error
	if err != nil {
		t.Error(err)
	} else if len(music.Artists) != 0 {
		t.Error("Artist was not removed from music")
	}
}

// GET FUNCTIONS
func TestMusicExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 32",
		Path:  "test_music_32.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	if !MusicExists(music.Id) {
		t.Error("Music does not exist")
	}
}

func TestMusicsExist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 33",
		Path:  "test_music_33.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	if !MusicsExist([]int{music.Id}) {
		t.Error("Music does not exist")
	}
}

func TestGetMusicJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 34",
		Path:  "test_music_34.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_json, err := GetMusicJSON(music.Id)
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(string(music_json), "Test Music 34") {
		t.Error("Music JSON is incorrect")
	}
}

func TestGetMusicsJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 35",
		Path:  "test_music_35.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	musics_json, err := GetMusicsJSON([]int{music.Id})
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(string(musics_json), "Test Music 35") {
		t.Error("Music JSON is incorrect")
	}
}

func TestGetMusicPathFromId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 36",
		Path:  "test_music_36.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	music_path, err := GetMusicPathFromId(music.Id)
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(music_path, "test_music_36.mp3") {
		t.Error("Music path is incorrect")
	}
}

func TestGetMusicsPathFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 37",
		Path:  "test_music_37.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	paths, err := GetMusicsPathFromIds([]int{music.Id})
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(paths[music.Id], "test_music_37.mp3") {
		t.Error("Music path is incorrect")
	}
}
