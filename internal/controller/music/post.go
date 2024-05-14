package music_controller

import (
	db_model "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
)

// Checks that all fields are valid, and posts the music to the database and saves the file to the server
// Returns an error if the music already exists or if there was an error saving the file to the server
// Returns nil if the music was successfully saved
func PostMusic(title string, genres []string, album_id int, music_file multipart.File, illustration_file_name string) (int, error) {
	music_file_name, err := file_utils.UploadMusicToServer(music_file)
	if err != nil {
		return -1, err
	}

	// Open the database, create the music, and close the database
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return music_model.CreateMusic(db, title, genres, album_id, music_file_name, illustration_file_name)
}
