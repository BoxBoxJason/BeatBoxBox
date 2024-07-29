package user_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"strings"
	"testing"
)

// POST FUNCTIONS
func TestPostUser(t *testing.T) {
	_, err := PostUser("Test User 31", "TestEmail31@email.com", "Password31_", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestAttemptLogin(t *testing.T) {
	user_id, err := PostUser("Test User 46", "TestEmail45@email.com", "Password45_", nil)
	if err != nil {
		t.Error(err)
	}
	user_id_login, _, err := AttemptLogin("Test User 46", "Password45_")
	if err != nil {
		t.Error(err)
	} else if user_id != user_id_login {
		t.Error("Wrong user logged in")
	}
}

// DELETE FUNCTIONS
func TestDeleteUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 32",
		Email:          "Test Email 32",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	user_id := user.Id
	err = DeleteUser(user_id)
	if err != nil {
		t.Error(err)
	}
	err = db.First(&db_tables.User{}, user_id).Error
	if err == nil {
		t.Error("User not deleted")
	}
}

func TestDeleteUsers(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 33",
		Email:          "Test Email 33",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	user_id := user.Id
	err = DeleteUsers([]int{user_id})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&db_tables.User{}, user_id).Error
	if err == nil {
		t.Error("User not deleted")
	}
}

// GET FUNCTIONS
func TestGetUserJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 34",
		Email:          "Test Email 34",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	user_json, err := GetUserJSON(user.Id)
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(string(user_json), "Test User 34") {
		t.Error("User not found in JSON")
	}
}

func TestGetUsersJSON(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 35",
		Email:          "Test Email 35",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	users_json, err := GetUsersJSON([]int{user.Id})
	if err != nil {
		t.Error(err)
	} else if !strings.Contains(string(users_json), "Test User 35") {
		t.Error("User not found in JSON")
	}
}

func TestGetUserIdFromUsername(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 36",
		Email:          "Test Email 36",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	user_id, err := GetUserIdFromUsername("Test User 36")
	if err != nil {
		t.Error(err)
	} else if user_id != user.Id {
		t.Errorf("User not found, got: %d instead of %d", user_id, user.Id)
	}
}

func TestUserExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 37",
		Email:          "Test Email 37",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	if !UserExists(user.Id) {
		t.Error("User not found")
	}
}

func TestUsersExist(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 38",
		Email:          "Test Email 38",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	if !UsersExist([]int{user.Id}) {
		t.Error("User not found")
	}
}

// PUT FUNCTIONS
func TestUpdateUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 39",
		Email:          "Test Email 39",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	user_id := user.Id
	err = UpdateUser(user_id, map[string]interface{}{"pseudo": "Test User 39 updated"})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&user, user_id).Error
	if err != nil {
		t.Error(err)
	} else if user.Pseudo != "Test User 39 updated" {
		t.Error("User not updated")
	}
}

func TestAddMusicsToLikedMusics(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 40",
		Email:          "Test Email 40",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	music := db_tables.Music{
		Title: "Test Music 43",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	err = AddMusicsToLikedMusics(user.Id, []int{music.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("LikedMusics").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.LikedMusics) == 0 {
		t.Error("Music not added to liked musics")
	}
}

func TestRemoveMusicsFromLikedMusics(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	music := db_tables.Music{
		Title: "Test Music 44",
		Path:  "test.mp3",
	}
	err = db.Create(&music).Error
	if err != nil {
		t.Error(err)
	}
	user := db_tables.User{
		Pseudo:         "Test User 41",
		Email:          "Test Email 41",
		HashedPassword: "hashed_password",
		LikedMusics:    []db_tables.Music{music},
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	err = RemoveMusicsFromLikedMusics(user.Id, []int{music.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("LikedMusics").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.LikedMusics) != 0 {
		t.Error("Music not removed from liked musics")
	}
}

func TestAddPlaylistsToOwnedPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 42",
		Email:          "Test Email 42",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title: "Test Playlist 26",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	err = AddPlaylistsToOwnedPlaylists(user.Id, []int{playlist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Playlists").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.Playlists) == 0 {
		t.Error("Playlist not added to owned playlists")
	}
}

func TestRemovePlaylistsFromOwnedPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 27",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	user := db_tables.User{
		Pseudo:         "Test User 43",
		Email:          "Test Email 43",
		HashedPassword: "hashed_password",
		Playlists:      []db_tables.Playlist{playlist},
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	err = RemovePlaylistsFromOwnedPlaylists(user.Id, []int{playlist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("Playlists").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.Playlists) != 0 {
		t.Error("Playlist not removed from owned playlists")
	}
}

func TestAddPlaylistsToSubscribedPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:         "Test User 44",
		Email:          "Test Email 44",
		HashedPassword: "hashed_password",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	playlist := db_tables.Playlist{
		Title: "Test Playlist 28",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	err = AddPlaylistsToSubscribedPlaylists(user.Id, []int{playlist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("SubscribedPlaylists").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.SubscribedPlaylists) == 0 {
		t.Error("Playlist not added to subscribed playlists")
	}
}

func TestRemovePlaylistsFromSubscribedPlaylists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	playlist := db_tables.Playlist{
		Title: "Test Playlist 29",
	}
	err = db.Create(&playlist).Error
	if err != nil {
		t.Error(err)
	}
	user := db_tables.User{
		Pseudo:              "Test User 45",
		Email:               "Test Email 45",
		HashedPassword:      "hashed_password",
		SubscribedPlaylists: []db_tables.Playlist{playlist},
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	err = RemovePlaylistsFromSubscribedPlaylists(user.Id, []int{playlist.Id})
	if err != nil {
		t.Error(err)
	}
	err = db.Preload("SubscribedPlaylists").First(&user, user.Id).Error
	if err != nil {
		t.Error(err)
	} else if len(user.SubscribedPlaylists) != 0 {
		t.Error("Playlist not removed from subscribed playlists")
	}
}
