package utils

import (
	"crypto/rand"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Return a 32 character long random string
func createRandomFileName(extension string) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes) + "." + extension, nil
}

// Create a filename that doesn't exist in the music directory
func CreateNonExistingMusicFileName() (string, error) {
	return createNonExistingFileName(filepath.Join("data", "musics"), "mp3")
}

// Create a filename that doesn't exist in the illustration directory
func CreateNonExistingIllustrationFileName() (string, error) {
	return createNonExistingFileName(filepath.Join("data", "illustrations"), "jpg")
}

// Create a file name that doesn't exist in the specified directory
func createNonExistingFileName(directory string, extension string) (string, error) {
	for {
		file_name, err := createRandomFileName(extension)
		if err != nil {
			return "", err
		}
		if _, err := os.Stat(filepath.Join(directory, file_name)); os.IsNotExist(err) {
			return file_name, nil
		}
	}
}

// Upload a file to the server
func UploadFileToServer(file multipart.File, dest_file string) error {
	out_file, err := os.Create(dest_file)
	if err != nil {
		return err
	}
	defer out_file.Close()

	_, err = io.Copy(out_file, file)
	if err != nil {
		return err
	}
	return nil
}
