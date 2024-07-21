package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

func PostPlaylist(title string, owners_ids []int, description string, illustration_file *multipart.File) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)
	if playlist_model.PlaylistAlreadyExists(db, title, owners_ids) {
		return -1, custom_errors.NewConflictError("playlist with same name & creator already exists")
	}
	owners, err := user_model.GetUsers(db, owners_ids)
	if err != nil || owners == nil || len(owners) != len(owners_ids) {
		return -1, custom_errors.NewNotFoundError("some users were not found")
	}
	owners_ptr := make([]*db_tables.User, len(owners))
	for i, owner := range owners {
		owners_ptr[i] = &owner
	}

	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "playlists")
	if err != nil {
		return -1, custom_errors.NewInternalServerError("Failed to upload illustration file")
	}
	return playlist_model.CreatePlaylist(db, title, owners_ptr, description, illustration_file_name)
}
