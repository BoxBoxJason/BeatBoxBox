package cookie_model

import (
	db_model "BeatBoxBox/internal/model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"testing"
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
	_, err = CreateCookie(db, hashed_token, 0)

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
	cookie := db_model.AuthCookie{
		HashedAuthToken: hashed_token,
	}
	db.Create(&cookie)
	cookie_id := cookie.Id
	UpdateCookieAuthToken(db, cookie.Id, "new_hashed_token")
	cookie = db_model.AuthCookie{}
	db.Where("id = ?", cookie_id).First(&cookie)

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
	cookie := db_model.AuthCookie{
		HashedAuthToken: hashed_token,
		UserId:          0,
	}
	db.Create(&cookie)

	cookies, err := GetUserCookies(db, 0)
	if err != nil {
		t.Errorf("Error getting cookie: %s", err)
	}
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
	cookie := db_model.AuthCookie{
		HashedAuthToken: hashed_token,
	}
	db.Create(&cookie)
	cookie_id := cookie.Id
	err = DeleteAuthToken(db, cookie.Id)
	if err != nil {
		t.Errorf("Error deleting cookie: %s", err)
	}

	cookie = db_model.AuthCookie{}
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
	cookie := db_model.AuthCookie{
		HashedAuthToken: hashed_token,
		ExpirationDate:  0,
	}
	db.Create(&cookie)
	cookie_id := cookie.Id
	err = DeleteExpiredTokens(db)
	if err != nil {
		t.Errorf("Error deleting expired tokens: %s", err)
	}

	cookie = db_model.AuthCookie{}
	result := db.Where("id = ?", cookie_id).First(&cookie)
	if result.Error == nil {
		t.Errorf("Expected cookie to be deleted, got %v", cookie)
	}
}
