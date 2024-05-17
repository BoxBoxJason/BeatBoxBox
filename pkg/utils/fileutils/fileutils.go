package file_utils

import (
	custom_errors "BeatBoxBox/pkg/errors"
	"archive/zip"
	"crypto/rand"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const MAX_MUSIC_FILE_SIZE = 25 * 1024 * 1024
const MAX_IMAGE_FILE_SIZE = 5 * 1024 * 1024
const MAX_REQUEST_SIZE = MAX_IMAGE_FILE_SIZE + MAX_MUSIC_FILE_SIZE + 1024
const DEFAULT_ILLUSTRATION_FILE = "default.jpg"

var PROJECT_ROOT_DIR string

func init() {
	PROJECT_ROOT_DIR, _ = os.Getwd()

}

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
func CreateNonExistingIllustrationFileName(illustration_directory string) (string, error) {
	return createNonExistingFileName(filepath.Join("data", "illustrations", illustration_directory), "jpg")
}

func UploadIllustrationToServer(illustration_header *multipart.FileHeader, illustration_file multipart.File, illustration_directory string) (string, error) {
	err := CheckFileMeetsRequirements(*illustration_header, MAX_IMAGE_FILE_SIZE, "image/jpeg")
	if err != nil {
		return DEFAULT_ILLUSTRATION_FILE, custom_errors.NewBadRequestError("Image does not meet requirements: " + err.Error())
	}
	illustration_file_name, err := CreateNonExistingIllustrationFileName(illustration_directory)
	if err != nil {
		return DEFAULT_ILLUSTRATION_FILE, err
	}
	err = UploadFileToServer(illustration_file, filepath.Join("data", "illustrations", illustration_directory, illustration_file_name))
	if err != nil {
		return DEFAULT_ILLUSTRATION_FILE, err
	}
	return illustration_file_name, nil
}

func UploadMusicToServer(music_file multipart.File) (string, error) {
	music_file_name, err := CreateNonExistingMusicFileName()
	if err != nil {
		return "", err
	}
	err = UploadFileToServer(music_file, filepath.Join("data", "musics", music_file_name))
	if err != nil {
		return "", err
	}
	return music_file_name, nil
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

// Checks if a directory path exists and creates it if it does not
func CheckDirExists(dir_path string) error {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		return os.MkdirAll(dir_path, 0755)
	}
	return nil
}

func CheckFileMeetsRequirements(file_header multipart.FileHeader, max_size int, content_type string) error {
	if file_header.Size > int64(max_size) {
		return errors.New("File too big, max size for " + content_type + " is " + strconv.Itoa(max_size/1024/1024) + "Mb (Megabytes)")
	}
	if file_header.Header.Get("Content-Type") != content_type {
		return errors.New("Invalid file type, should be " + content_type)
	}
	return nil
}

// Zip files into a single zip file and write it to the response writer
// Takes a list of file paths and writes them to the zip file
func ServeZip(w http.ResponseWriter, files_paths []string, zip_file_name string) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+zip_file_name+".zip")

	zip_writer := zip.NewWriter(w)
	defer zip_writer.Close()

	for _, file_path := range files_paths {
		file, err := os.Open(file_path)
		if err != nil {
			http.Error(w, "File not found: "+file_path, http.StatusInternalServerError)
			return
		}
		defer file.Close()
		zip_file, err := zip_writer.Create(filepath.Base(file_path))
		if err != nil {
			http.Error(w, "Error creating zip file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(zip_file, file)
		if err != nil {
			http.Error(w, "Error copying file to zip: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func ServeTreeZip(w http.ResponseWriter, files_paths map[string][]string, zip_file_name string) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+zip_file_name+".zip")

	zip_writer := zip.NewWriter(w)
	defer zip_writer.Close()

	for directory_name, files_paths := range files_paths {
		_, err := zip_writer.Create(directory_name + "/")
		if err != nil {
			http.Error(w, "Error creating zip folder: "+err.Error(), http.StatusInternalServerError)
			return
		}
		for _, file_path := range files_paths {
			file, err := os.Open(file_path)
			if err != nil {
				http.Error(w, "File not found: "+file_path, http.StatusInternalServerError)
				return
			}
			defer file.Close()
			zip_file, err := zip_writer.Create(filepath.Join(directory_name, filepath.Base(file_path)))
			if err != nil {
				http.Error(w, "Error creating zip file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.Copy(zip_file, file)
			if err != nil {
				http.Error(w, "Error copying file to zip: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
