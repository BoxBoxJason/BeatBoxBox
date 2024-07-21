package user_controller

import (
	db_tables "BeatBoxBox/internal/model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"encoding/json"
	"path/filepath"
)

func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "users"))
}

func ConvertUserToJSON(user *db_tables.User) ([]byte, error) {
	subscribed_playlists_ids := make([]int, len(user.SubscribedPlaylists))
	for i, playlist := range user.SubscribedPlaylists {
		subscribed_playlists_ids[i] = playlist.Id
	}
	playlists_ids := make([]int, len(user.Playlists))
	for i, playlist := range user.Playlists {
		playlists_ids[i] = playlist.Id
	}
	liked_musics_ids := make([]int, len(user.LikedMusics))
	for i, music := range user.LikedMusics {
		liked_musics_ids[i] = music.Id
	}

	user_json := map[string]interface{}{
		"id":                   user.Id,
		"pseudo":               user.Pseudo,
		"email":                user.Email,
		"illustration":         user.Illustration,
		"playlists":            playlists_ids,
		"liked_musics":         liked_musics_ids,
		"subscribed_playlists": subscribed_playlists_ids,
		"uploaded_musics":      user.UploadedMusics,
		"created_on":           user.CreatedOn,
		"modified_on":          user.ModifiedOn,
	}
	return json.Marshal(user_json)
}

func ConvertUsersToJSON(users []*db_tables.User) ([]byte, error) {
	users_json := make([]map[string]interface{}, len(users))
	for i, user := range users {
		subscribed_playlists_ids := make([]int, len(user.SubscribedPlaylists))
		for i, playlist := range user.SubscribedPlaylists {
			subscribed_playlists_ids[i] = playlist.Id
		}
		playlists_ids := make([]int, len(user.Playlists))
		for i, playlist := range user.Playlists {
			playlists_ids[i] = playlist.Id
		}
		liked_musics_ids := make([]int, len(user.LikedMusics))
		for i, music := range user.LikedMusics {
			liked_musics_ids[i] = music.Id
		}
		users_json[i] = map[string]interface{}{
			"id":                   user.Id,
			"pseudo":               user.Pseudo,
			"email":                user.Email,
			"illustration":         user.Illustration,
			"playlists":            playlists_ids,
			"liked_musics":         liked_musics_ids,
			"subscribed_playlists": subscribed_playlists_ids,
			"uploaded_musics":      user.UploadedMusics,
			"created_on":           user.CreatedOn,
			"modified_on":          user.ModifiedOn,
		}
	}
	return json.Marshal(users_json)
}
