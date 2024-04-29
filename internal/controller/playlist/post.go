package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	"BeatBoxBox/pkg/utils"
	"errors"
	"mime/multipart"
	"path/filepath"
)

func PostPlaylist(title string, creator_id int, illustration_file multipart.File) error {
	// Check if the title is empty or already taken
	if title == "" {
		return errors.New("playlist title is empty")
	}
	if PlaylistExistsFromParams(title, creator_id) {
		return errors.New("playlist with same name & creator already exists")
	}

	// Generate a new file name & save the illustration file if needed
	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err := utils.CreateNonExistingIllustrationFileName("playlists")
		if err != nil {
			return err
		}
		err = utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", "playlists", illustration_file_name))
		if err != nil {
			return err
		}
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)

	return playlist_model.CreatePlaylist(db, title, creator_id, illustration_file_name)
}
