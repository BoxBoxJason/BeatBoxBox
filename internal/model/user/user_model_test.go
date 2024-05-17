package user_model

import (
	db_model "BeatBoxBox/internal/model"
	"testing"
)

// POST FUNCTIONS TESTS

func TestUserCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, err = CreateUser(db, "test user", "test email", "hashed_password")
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

	user := db_model.User{
		Pseudo: "test user",
	}
	db.Create(&user)
	user_id := user.Id

	UpdateUser(db, user_id, map[string]interface{}{"pseudo": "new test user"})
	user = db_model.User{}
	db.Where("id = ?", user_id).First(&user)

	if user.Pseudo != "new test user" {
		t.Errorf("Expected user pseudo to be 'new test user', got '%s'", user.Pseudo)
	}
}

// GET FUNCTIONS TESTS

func TestUserGet(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_model.User{
		Pseudo: "test user",
	}
	db.Create(&user)
	user_id := user.Id

	user, err = GetUser(db, user_id)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.Pseudo != "test user" {
		t.Errorf("Expected user pseudo to be 'test user', got '%s'", user.Pseudo)
	}
}

func TestUsersGet(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_model.User{
		Pseudo: "test user 1",
	}
	db.Create(&user1)
	user1_id := user1.Id

	user2 := db_model.User{
		Pseudo: "test user 2",
	}
	db.Create(&user2)
	user2_id := user2.Id

	users, err := GetUsers(db, []int{user1_id, user2_id})
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	}

	if len(users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(users))
	}
}

func TestUsersGetFromFilters(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user1 := db_model.User{
		Pseudo: "test user 1",
	}
	db.Create(&user1)

	user2 := db_model.User{
		Pseudo: "test user 2",
	}
	db.Create(&user2)

	users, err := GetUsersFromPartialPseudo(db, "test user")
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	}

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

	user1 := db_model.User{
		Pseudo: "test user 1",
	}
	db.Create(&user1)

	user2 := db_model.User{
		Pseudo: "test user 2",
	}
	db.Create(&user2)

	users, err := GetUsersFromPartialPseudo(db, "test user")
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	}

	if len(users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(users))
	}
}

// DELETE FUNCTIONS TESTS

func TestUserDelete(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	user := db_model.User{
		Pseudo: "test user",
	}
	db.Create(&user)
	user_id := user.Id

	err = DeleteUser(db, user_id)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	user = db_model.User{}
	result := db.Where("id = ?", user_id).First(&user)
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

	user1 := db_model.User{
		Pseudo: "test user 1",
	}
	db.Create(&user1)
	user1_id := user1.Id

	user2 := db_model.User{
		Pseudo: "test user 2",
	}
	db.Create(&user2)
	user2_id := user2.Id

	err = DeleteUsers(db, []int{user1_id, user2_id})
	if err != nil {
		t.Errorf("Error deleting users: %s", err)
	}

	users := []db_model.User{}
	result := db.Where("id IN ?", []int{user1_id, user2_id}).Find(&users)
	if result.Error == nil {
		t.Errorf("Expected an error, got nil")
	}
}
