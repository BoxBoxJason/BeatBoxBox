package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
)

func PostPlaylist(title string, description string, public bool, owners_ids []int, illustration_file *multipart.FileHeader, musics_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	if playlist_model.PlaylistAlreadyExists(db, title, owners_ids) {
		return []byte{}, httputils.NewConflictError("playlist with same name & creator already exists")
	}
	owners, err := user_model.GetUsers(db, owners_ids)
	if err != nil || owners == nil || len(owners) != len(owners_ids) {
		return []byte{}, httputils.NewNotFoundError("some users were not found")
	}
	owners_ptr := make([]*db_tables.User, len(owners))
	for i, owner := range owners {
		owners_ptr[i] = &owner
	}
	musics, err := music_model.GetMusics(db, musics_ids)
	if err != nil || len(musics) != len(musics_ids) {
		return []byte{}, httputils.NewNotFoundError("some musics were not found")
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "playlists")
	if err != nil {
		return []byte{}, httputils.NewInternalServerError("Failed to upload illustration file")
	}
	playlist, err := playlist_model.CreatePlaylist(db.Preload("Musics"), title, owners_ptr, description, public, illustration_file_name, musics_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}
