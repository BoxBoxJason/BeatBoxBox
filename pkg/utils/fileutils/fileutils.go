package file_utils

import (
	"BeatBoxBox/pkg/logger"
	bool_utils "BeatBoxBox/pkg/utils/boolutils"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const DEFAULT_ILLUSTRATION_FILE = "default.jpg"

var MUSICS_DIR string
var ILLUSTRATIONS_DIR string
var ILLUSTRATIONS_DIRS map[string]string

func init() {
	BEATBOXBOX_ROOT_DIR := os.Getenv("BEATBOXBOX_ROOT_DIR")
	if BEATBOXBOX_ROOT_DIR == "" {
		BEATBOXBOX_ROOT_DIR = "/home/user/beatboxbox"
	}
	MUSICS_DIR = filepath.Join(BEATBOXBOX_ROOT_DIR, "data", "musics")
	go CheckDirExists(MUSICS_DIR)

	ILLUSTRATIONS_DIR = filepath.Join(BEATBOXBOX_ROOT_DIR, "data", "illustrations")

	ILLUSTRATIONS_DIRS = map[string]string{
		"albums":    filepath.Join(ILLUSTRATIONS_DIR, "albums"),
		"artists":   filepath.Join(ILLUSTRATIONS_DIR, "artists"),
		"users":     filepath.Join(ILLUSTRATIONS_DIR, "users"),
		"musics":    filepath.Join(ILLUSTRATIONS_DIR, "musics"),
		"playlists": filepath.Join(ILLUSTRATIONS_DIR, "playlists"),
	}
	for _, dir := range ILLUSTRATIONS_DIRS {
		go CheckDirExists(dir)
	}
}

func createRandomFileName(extension string) string {
	const CHARACTERS = "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, 32)
	for i := range bytes {
		bytes[i] = CHARACTERS[rand.Intn(len(CHARACTERS))]
	}
	return string(bytes) + "." + extension
}

func createNonExistingMusicFileName(extension string) (string, error) {
	music_subdir, err := getLastSubdirectory(MUSICS_DIR)
	if err != nil {
		return "", err
	}
	new_music_file_name := createNonExistingFileName(filepath.Join(MUSICS_DIR, music_subdir), extension)
	return filepath.Join(music_subdir, new_music_file_name), nil
}

func getLastSubdirectory(directory_path string) (string, error) {
	number_dirs, err := countSubDirs(directory_path)
	if err != nil {
		return "", err
	}

	// Check if the last directory is full
	attempt_new_dir := strconv.Itoa(bool_utils.Max(number_dirs-1, 0))
	subdirs_count, err := countSubDirs(filepath.Join(directory_path, attempt_new_dir))
	if err != nil {
		return "", err
	}
	// If the last directory is full, create a new one
	if subdirs_count >= 1000 {
		attempt_new_dir = strconv.Itoa(number_dirs)
	}
	err = CheckDirExists(filepath.Join(directory_path, attempt_new_dir))
	if err != nil {
		return "", err
	}

	return attempt_new_dir, nil
}

func countSubDirs(directory string) (int, error) {
	sub_dirs, err := os.ReadDir(directory)
	return len(sub_dirs), err
}

func createNonExistingIllustrationFileName(illustration_directory string) (string, error) {
	illustration_subdir, err := getLastSubdirectory(ILLUSTRATIONS_DIRS[illustration_directory])
	if err != nil {
		return "", err
	}
	new_illustration_file_name := createNonExistingFileName(filepath.Join(ILLUSTRATIONS_DIRS[illustration_directory], illustration_subdir), "jpg")
	return filepath.Join(illustration_subdir, new_illustration_file_name), nil
}

func UploadIllustrationToServer(illustration_file *multipart.FileHeader, illustration_directory string) (string, error) {
	if illustration_file == nil {
		return DEFAULT_ILLUSTRATION_FILE, nil
	}
	illustration_file_name, err := createNonExistingIllustrationFileName(illustration_directory)
	if err != nil {
		return DEFAULT_ILLUSTRATION_FILE, err
	}
	err = UploadFileToServer(illustration_file, filepath.Join(ILLUSTRATIONS_DIRS[illustration_directory], illustration_file_name))
	if err != nil {
		return DEFAULT_ILLUSTRATION_FILE, err
	}
	return illustration_file_name, nil
}

func UploadMusicToServer(music_file *multipart.FileHeader) (string, error) {
	if music_file == nil {
		return "none", nil
	}
	extension := filepath.Ext(music_file.Filename)[1:]
	music_file_name, err := createNonExistingMusicFileName(extension)
	if err != nil {
		return "", err
	}
	err = UploadFileToServer(music_file, filepath.Join(MUSICS_DIR, music_file_name))
	if err != nil {
		return "", err
	}
	return music_file_name, nil
}

func createNonExistingFileName(directory string, extension string) string {
	for {
		file_name := createRandomFileName(extension)
		if _, err := os.Stat(filepath.Join(directory, file_name)); os.IsNotExist(err) {
			return file_name
		}
	}
}

func UploadFileToServer(file *multipart.FileHeader, dest_file string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			logger.Error("Error closing file: ", err)
		}
	}(src)

	out, err := os.Create(dest_file)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			logger.Error("Error closing file: ", err)
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}

func CheckDirExists(dir_path string) error {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		return os.MkdirAll(dir_path, 0755)
	}
	return nil
}

func ServeZip(w http.ResponseWriter, files_paths []string, zip_file_name string) {

}

func ServeTreeZip(w http.ResponseWriter, files_paths map[string][]string, zip_file_name string) {

}
