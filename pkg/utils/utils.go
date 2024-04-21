package utils

import (
	"crypto/rand"
	"os"
)

// Return a 32 character long random string
func createRandomFileName() (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes) + ".mp3", nil
}

// Create a filename that doesn't exist in the music directory
func CreateNonExistingMusicFileName() (string, error) {
	for {
		file_name, err := createRandomFileName()
		if err != nil {
			return "", err
		}
		if _, err := os.Stat("./data/musics/" + file_name); os.IsNotExist(err) {
			return file_name, nil
		}
	}
}
