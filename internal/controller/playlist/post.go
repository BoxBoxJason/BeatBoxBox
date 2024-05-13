package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"errors"
	"mime/multipart"
	"path/filepath"
)

func PostPlaylist(title string, creator_id int, description string, illustration_file multipart.File) (int, error) {
	// Check if the title is empty or already taken
	if title == "" {
		return -1, errors.New("playlist title is empty")
	}
	if PlaylistExistsFromParams(title, creator_id) {
		return -1, errors.New("playlist with same name & creator already exists")
	}

	// Generate a new file name & save the illustration file if needed
	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err := file_utils.CreateNonExistingIllustrationFileName("playlists")
		if err != nil {
			return -1, err
		}
		err = file_utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", "playlists", illustration_file_name))
		if err != nil {
			return -1, err
		}
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return playlist_model.CreatePlaylist(db, title, creator_id, description, illustration_file_name)
}
