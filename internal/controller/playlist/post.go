package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	"errors"
)

func PostPlaylist(title string, creator_id int, description string, illustration_file_name string) (int, error) {
	// Check if the title is empty or already taken
	if title == "" {
		return -1, errors.New("playlist title is empty")
	}
	if PlaylistExistsFromParams(title, creator_id) {
		return -1, errors.New("playlist with same name & creator already exists")
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return playlist_model.CreatePlaylist(db, title, creator_id, description, illustration_file_name)
}
