package user_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS TESTS

func TestUserCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreateUser(db, "Test user 1", "Test email 1", "hashed_password", "default.jpg")
	if err != nil {
		t.Errorf("Error creating user: %s", err)
	}
}

// PUT FUNCTIONS TESTS

func TestUserUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test user 2",
		Email:           "Test email 2",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	err = UpdateUser(db, &user, map[string]interface{}{"pseudo": "Renamed Test User 2"})
	if err != nil {
		t.Errorf("Error updating user: %s", err)
	}

	if user.Pseudo != "Renamed Test User 2" {
		t.Errorf("Expected user pseudo to be 'Renamed Test User 2', got '%s'", user.Pseudo)
	}
}

func TestUserAddSubscribedPlaylistsToUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 14",
		Email:           "Test email 14",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	playlist := db_tables.Playlist{
		Title: "Test Playlist 1",
	}
	db.Create(&playlist)

	err = AddSubscribedPlaylistsToUser(db, &user, []*db_tables.Playlist{&playlist})
	if err != nil {
		t.Errorf("Error adding playlist to user: %s", err)
	}
}

func TestUserRemoveSubscribedPlaylistsFromUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 15",
		Email:           "Test email 15",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	playlist := db_tables.Playlist{
		Title: "Test Playlist 2",
	}
	db.Create(&playlist)

	user.SubscribedPlaylists = append(user.SubscribedPlaylists, playlist)

	err = RemoveSubscribedPlaylistsFromUser(db, &user, []*db_tables.Playlist{&playlist})
	if err != nil {
		t.Errorf("Error removing playlist from user: %s", err)
	}

}

func TestUserAddLikedMusicsToUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 16",
		Email:           "Test email 16",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	music := db_tables.Music{
		Title: "Test Music 13",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	err = AddLikedMusicsToUser(db, &user, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error adding music to user: %s", err)
	}
}

func TestUserRemoveLikedMusicFromUser(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 17",
		Email:           "Test email 17",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	music := db_tables.Music{
		Title: "Test Music 14",
		Path:  "fake.mp3",
	}
	db.Create(&music)

	user.LikedMusics = append(user.LikedMusics, music)

	err = RemoveLikedMusicsFromUser(db, &user, []*db_tables.Music{&music})
	if err != nil {
		t.Errorf("Error removing music from user: %s", err)
	}
}

// GET FUNCTIONS TESTS

func TestUserGet(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 3",
		Email:           "Test email 3",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)
	user_id := user.Id

	user, err = GetUser(db, user_id)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.Pseudo != "Test User 3" {
		t.Errorf("Expected user pseudo to be 'Test User 3', got '%s'", user.Pseudo)
	}
}

func TestUsersGet(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_tables.User{
		Pseudo:          "Test User 4",
		Email:           "Test email 4",
		Hashed_password: "hashed_password",
	}
	db.Create(&user1)
	user1_id := user1.Id

	user2 := db_tables.User{
		Pseudo:          "Test User 5",
		Email:           "Test email 5",
		Hashed_password: "hashed_password",
	}
	db.Create(&user2)
	user2_id := user2.Id

	users, err := GetUsers(db, []int{user1_id, user2_id})
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	} else if len(users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(users))
	}
}

func TestUsersGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_tables.User{
		Pseudo:          "Test User 6",
		Email:           "Test email 6",
		Hashed_password: "hashed_password",
	}
	db.Create(&user1)

	user2 := db_tables.User{
		Pseudo:          "Test User 7",
		Email:           "Test email 7",
		Hashed_password: "hashed_password",
	}
	db.Create(&user2)

	users := GetUsersFromFilters(db, map[string]interface{}{"hashed_password": "hashed_password"})
	if len(users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(users))
	}
}

func TestUserGetFromPartialName(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_tables.User{
		Pseudo:          "Test User 8",
		Email:           "Test email 8",
		Hashed_password: "hashed_password",
	}
	db.Create(&user1)

	user2 := db_tables.User{
		Pseudo:          "Test User 9",
		Email:           "Test email 9",
		Hashed_password: "hashed_password",
	}
	db.Create(&user2)

	users, err := GetUsersFromPartialPseudo(db, "Test User")
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	} else if len(users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(users))
	}
}

func TestUserAlreadyExists(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 30",
		Email:           "Test email 30",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	exists, fields, err := UserAlreadyExists(db, "Test User 30", "")
	if err != nil {
		t.Errorf("Error checking if user already exists: %s", err)
	} else if !exists || len(fields) != 1 || fields[0] != "pseudo" {
		t.Errorf("Expected user to exist, got false or wrong fields")
	}

	exists, fields, err = UserAlreadyExists(db, "USERNOTEXIST", "USERNOTEXIST")
	if err != nil {
		t.Errorf("Error checking if user already exists: %s", err)
	} else if exists {
		t.Errorf("Expected user to not exist, got true")
	}
}

// DELETE FUNCTIONS TESTS

func TestUserDelete(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 10",
		Email:           "Test email 10",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)
	user_id := user.Id

	err = DeleteUser(db, user_id)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	user = db_tables.User{}
	result := db.Where("id = ?", user_id).First(&user)
	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestUserDeleteFromRecord(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_tables.User{
		Pseudo:          "Test User 11",
		Email:           "Test email 11",
		Hashed_password: "hashed_password",
	}
	db.Create(&user)

	err = DeleteUserFromRecord(db, &user)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	user = db_tables.User{}
	result := db.Where("id = ?", user.Id).First(&user)
	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestUsersDeleteFromIds(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_tables.User{
		Pseudo:          "Test User 12",
		Email:           "Test email 12",
		Hashed_password: "hashed_password",
	}
	db.Create(&user1)
	user_id := user1.Id

	err = DeleteUsers(db, []int{user_id})
	if err != nil {
		t.Errorf("Error deleting users: %s", err)
	} else {
		user1 = db_tables.User{}
		result := db.Where("id = ?", user_id).First(&user1)
		if result.Error == nil {
			t.Errorf("Expected an error, got nil")
		}
	}
}

func TestUsersDeleteFromRecords(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_tables.User{
		Pseudo:          "Test User 13",
		Email:           "Test email 13",
		Hashed_password: "hashed_password",
	}
	db.Create(&user1)

	user_id := user1.Id
	err = DeleteUsersFromRecords(db, []*db_tables.User{&user1})
	if err != nil {
		t.Errorf("Error deleting users: %s", err)
	} else {
		user1 = db_tables.User{}
		result := db.Where("id = ?", user_id).First(&user1)
		if result.Error == nil {
			t.Errorf("Expected an error, got nil")
		}
	}
}
