package cookie_controller

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"testing"
	"time"
)

// POST FUNCTIONS
func TestPostAuthToken(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo: "Test User 23",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	_, err = PostAuthToken(user.Id)
	if err != nil {
		t.Error(err)
	}
}

// DELETE FUNCTIONS
func TestDeleteMatchingAuthToken(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 24",
		Hashed_password: "password",
		Email:           "Test email 24",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	authToken, err := PostAuthToken(user.Id)
	if err != nil {
		t.Error(err)
	}
	err = DeleteMatchingAuthToken(user.Id, authToken)
	if err != nil {
		t.Error(err)
	}
}

// GET FUNCTIONS
func TestCheckAuthTokenMatches(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 25",
		Hashed_password: "password",
		Email:           "Test email 25",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	auth_cookie := db_tables.AuthCookie{
		UserId:          user.Id,
		ExpirationDate:  auth_utils.GetNewTokenExpirationTime(),
		HashedAuthToken: "hashed_auth_token",
	}
	err = db.Create(&auth_cookie).Error
	if err != nil {
		t.Error(err)
	}
	_, _, err = CheckAuthTokenMatches(user.Id, auth_cookie.HashedAuthToken)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMatchingAuthTokenId(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 26",
		Hashed_password: "password",
		Email:           "Test email 26",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}

	auth_cookie := db_tables.AuthCookie{
		UserId:          user.Id,
		ExpirationDate:  auth_utils.GetNewTokenExpirationTime(),
		HashedAuthToken: "hashed_auth_token",
	}
	err = db.Create(&auth_cookie).Error
	if err != nil {
		t.Error(err)
	}

	_, err = GetMatchingAuthTokenId(user.Id, auth_cookie.HashedAuthToken)
	if err != nil {
		t.Error(err)
	}
}

// PUT FUNCTIONS
func TestupdateAuthTokenIfNearExpiry(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	user := db_tables.User{
		Pseudo:          "Test User 27",
		Hashed_password: "password",
		Email:           "Test email 27",
	}
	err = db.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
	auth_cookie := db_tables.AuthCookie{
		UserId:          user.Id,
		ExpirationDate:  time.Now().Unix(),
		HashedAuthToken: "hashed_auth_token",
	}
	err = db.Create(&auth_cookie).Error
	if err != nil {
		t.Error(err)
	}
	new_token, err := updateAuthTokenIfNearExpiry(&auth_cookie)
	if err != nil {
		t.Error(err)
	} else if new_token == "" {
		t.Error("new token not generated")
	}
	result := db.Where("id = ?", auth_cookie.Id).First(&auth_cookie)
	if result.Error != nil {
		t.Error(result.Error)
	} else if auth_cookie.HashedAuthToken == "hashed_auth_token" {
		t.Error("auth token not updated")
	}
}
