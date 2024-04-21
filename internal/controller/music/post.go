package music_controller

import (
	music_model "BeatBoxBox/internal/model/music"
	"BeatBoxBox/pkg/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Checks that all fields are valid, and posts the music to the database and saves the file to the server
// Returns an error if the music already exists or if there was an error saving the file to the server
// Returns nil if the music was successfully saved
func PostMusic(title string, author string, genres []string, album string, file multipart.File) error {
	// Generate a new file name
	file_name, err := utils.CreateNonExistingMusicFileName()
	if err != nil {
		return err
	}
	dest_file := filepath.Join("data", "musics", file_name)

	// Create the file
	out_file, err := os.Create(dest_file)
	if err != nil {
		return err
	}
	defer out_file.Close()

	// Copy the file to the server
	_, err = io.Copy(out_file, file)
	if err != nil {
		return err
	}

	if musicExists(title) {
		return fmt.Errorf("Music already exists")
	}

	if err := music_model.CreateMusic(title, author, genres, album, file_name); err != nil {
		return err
	}
	return nil
}
