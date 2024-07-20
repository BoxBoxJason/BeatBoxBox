package playlist_controller

import (
	playlist_model "BeatBoxBox/internal/model/playlist"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

func PostPlaylist(title string, creator_id int, description string, illustration_file *multipart.File) (int, error) {
	if PlaylistAlreadyExists(title, creator_id) {
		return -1, custom_errors.NewConflictError("playlist with same name & creator already exists")
	}

	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "playlists")
	if err != nil {
		return -1, custom_errors.NewInternalServerError("Failed to upload illustration file")
	}
	return playlist_model.CreatePlaylist(db, title, creator_id, description, illustration_file_name)
}
