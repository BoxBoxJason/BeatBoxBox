package user_controller

import (
	db_model "BeatBoxBox/internal/model"
	user_model "BeatBoxBox/internal/model/user"
	"encoding/json"
)

// UserExists returns whether a user exists in the database
func UserExists(user_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = user_model.GetUser(db, user_id)
	return err == nil
}

// UsersExist returns whether a list of users exists in the database
func UsersExist(user_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsers(db, user_ids)
	return err == nil && len(users) == len(user_ids)
}

// UserExistsFromParams returns whether a user exists in the database
func UserExistsFromName(username string) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{
		"Pseudo": username,
	})
	return err == nil && len(users) > 0
}

// UserExistsFromEmail returns whether a user exists in the database
func UserExistsFromEmail(email string) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{
		"Email": email,
	})
	return err == nil && len(users) > 0
}

// GetUser returns a user from the database
// Selects the user with the given user_id
func GetUser(user_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return nil, err
	}

	return mapUserToJson(user)
}

// GetUsers returns a list of users from the database
// Selects the users with the given user_ids
func GetUsers(user_ids []int) ([][]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	users, err := user_model.GetUsers(db, user_ids)
	if err != nil {
		return nil, err
	}

	users_json := make([][]byte, len(users))
	for i, user := range users {
		user_json, err := mapUserToJson(user)
		if err != nil {
			return nil, err
		}
		users_json[i] = user_json
	}
	return users_json, nil
}

func GetUserIdFromUsername(username string) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return 0, err
	}
	defer db_model.CloseDB(db)
	users, err := user_model.GetUsersFromFilters(db, map[string]interface{}{
		"Pseudo": username,
	})
	if err != nil || len(users) == 0 {
		return 0, err
	}
	return users[0].Id, nil
}

func mapUserToJson(user db_model.User) ([]byte, error) {
	playlists_ids := make([]int, len(user.Playlists))
	for i, playlist := range user.Playlists {
		playlists_ids[i] = playlist.Id
	}

	subscribed_playlists_ids := make([]int, len(user.SubscribedPlaylists))
	for i, playlist := range user.SubscribedPlaylists {
		subscribed_playlists_ids[i] = playlist.Id
	}

	uploaded_musics_ids := make([]int, len(user.UploadedMusics))
	for i, music := range user.UploadedMusics {
		uploaded_musics_ids[i] = music.Id
	}

	return json.Marshal(map[string]interface{}{
		"id":                   user.Id,
		"pseudo":               user.Pseudo,
		"illustration":         user.Illustration,
		"subscribed_playlists": subscribed_playlists_ids,
		"playlists":            playlists_ids,
		"uploaded_musics":      uploaded_musics_ids,
	})
}
