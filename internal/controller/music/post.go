package music_controller

import (
	db_model "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"mime/multipart"
	"path/filepath"
)

// Checks that all fields are valid, and posts the music to the database and saves the file to the server
// Returns an error if the music already exists or if there was an error saving the file to the server
// Returns nil if the music was successfully saved
func PostMusic(title string, artist_id int, genres []string, album_id int, music_file multipart.File, illustration_file multipart.File) (int, error) {
	// Generate a new file name & save the music file
	music_file_name, err := file_utils.CreateNonExistingMusicFileName()
	if err != nil {
		return -1, err
	}
	err = file_utils.UploadFileToServer(music_file, filepath.Join("data", "musics", music_file_name))
	if err != nil {
		return -1, err
	}

	// Generate a new file name & save the illustration file
	illustration_file_name := "default.jpg"
	if illustration_file != nil {
		illustration_file_name, err = file_utils.CreateNonExistingIllustrationFileName("musics")
		if err != nil {
			return -1, err
		}
		err = file_utils.UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", "musics", illustration_file_name))
		if err != nil {
			return -1, err
		}
	}

	// Open the database, create the music, and close the database
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	return music_model.CreateMusic(db, title, artist_id, genres, album_id, music_file_name, illustration_file_name)
}
