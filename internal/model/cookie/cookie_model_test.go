package cookie_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"testing"
	"time"
)

// POST FUNCTIONS TESTS
func TestCookieCreate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, hashed_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		t.Errorf("Error generating random token: %s", err)
	}
	user := db_tables.User{
		Pseudo:          "Test User 18",
		Hashed_password: "password",
		Email:           "Test email 18",
	}
	db.Create(&user)

	_, err = CreateCookie(db, hashed_token, user.Id)

	if err != nil {
		t.Errorf("Error creating cookie: %s", err)
	}
}

// PUT FUNCTIONS TESTS
func TestCookieUpdate(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, hashed_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		t.Errorf("Error generating random token: %s", err)
	}

	user := db_tables.User{
		Pseudo:          "Test User 19",
		Hashed_password: "password",
		Email:           "Test email 19",
	}
	db.Create(&user)

	cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		UserId:          user.Id,
		ExpirationDate:  time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix(),
	}
	db.Create(&cookie)
	err = UpdateCookieAuthToken(db, &cookie, "new_hashed_token")

	if cookie.HashedAuthToken != "new_hashed_token" {
		t.Errorf("Expected cookie hashed token to be 'new_hashed_token', got '%s'", cookie.HashedAuthToken)
	}
}

// GET FUNCTIONS TESTS
func TestGetUserCookie(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, hashed_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		t.Errorf("Error generating random token: %s", err)
	}

	user := db_tables.User{
		Pseudo:          "Test User 20",
		Hashed_password: "password",
		Email:           "Test email 20",
	}
	db.Create(&user)

	cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		UserId:          user.Id,
		ExpirationDate:  time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix(),
	}
	db.Create(&cookie)

	cookies := GetUserCookies(db, user.Id)
	if len(cookies) < 1 {
		t.Errorf("Expected at least 1 cookie, got %d", len(cookies))
	}
}

// DELETE FUNCTIONS TESTS
func TestCookieDeleteAuthToken(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, hashed_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		t.Errorf("Error generating random token: %s", err)
	}

	user := db_tables.User{
		Pseudo:          "Test User 21",
		Hashed_password: "password",
		Email:           "Test email 21",
	}
	db.Create(&user)

	cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		ExpirationDate:  time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION).Unix(),
		UserId:          user.Id,
	}
	db.Create(&cookie)
	cookie_id := cookie.Id
	err = DeleteAuthToken(db, &cookie)
	if err != nil {
		t.Errorf("Error deleting cookie: %s", err)
	}

	cookie = db_tables.AuthCookie{}
	result := db.Where("id = ?", cookie_id).First(&cookie)
	if result.Error == nil {
		t.Errorf("Expected cookie to be deleted, got %v", cookie)
	}
}

func TestCookieDeleteExpiredTokens(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Errorf("Error opening database: %s", err)
	}
	defer db_model.CloseDB(db)

	_, hashed_token, err := auth_utils.GenerateRandomTokenWithHash()
	if err != nil {
		t.Errorf("Error generating random token: %s", err)
	}

	user := db_tables.User{
		Pseudo:          "Test User 22",
		Hashed_password: "password",
		Email:           "Test email 22",
	}
	db.Create(&user)

	cookie := db_tables.AuthCookie{
		HashedAuthToken: hashed_token,
		ExpirationDate:  0,
		UserId:          user.Id,
	}
	db.Create(&cookie)
	cookie_id := cookie.Id
	err = DeleteExpiredTokens(db)
	if err != nil {
		t.Errorf("Error deleting expired tokens: %s", err)
	}

	cookie = db_tables.AuthCookie{}
	result := db.Where("id = ?", cookie_id).First(&cookie)
	if result.Error == nil {
		t.Errorf("Expected cookie to be deleted, got %v", cookie)
	}
}
