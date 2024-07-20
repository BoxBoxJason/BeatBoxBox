package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

// Checks that all fields are valid, and posts the music to the database and saves the file to the server
func PostMusic(title string, genres []string, lyrics string, album_id int, music_file *multipart.File, illustration_file *multipart.File, artists_ids []int) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	if music_model.MusicAlreadyExists(db, title, artists_ids) {
		return -1, custom_errors.NewConflictError("music already exists")
	}

	artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return -1, err
	} else if artists == nil || len(artists) != len(artists_ids) {
		return -1, custom_errors.NewNotFoundError("some artists were not found")
	}
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}

	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "musics")
	if err != nil {
		return -1, custom_errors.NewInternalServerError("could not upload illustration")
	}
	music_file_name, err := file_utils.UploadMusicToServer(music_file)
	if err != nil {
		return -1, custom_errors.NewInternalServerError("could not upload music")
	}

	return music_model.CreateMusic(db, title, genres, lyrics, album_id, music_file_name, illustration_file_name, artists_ptr)
}
