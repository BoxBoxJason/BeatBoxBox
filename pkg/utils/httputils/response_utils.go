package httputils

import (
	"BeatBoxBox/pkg/logger"
	auth_utils "BeatBoxBox/pkg/utils/authutils"
	"archive/zip"
	"net/http"
	"os"
	"time"
)

// RespondWithJSON sends a JSON response to the client with the given status code and content
func RespondWithJSON(w http.ResponseWriter, status int, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(content)
	if err != nil {
		SendErrorToClient(w, NewInternalServerError("Error writing response: "+err.Error()))
	}
}

// ServeZip creates a zip file from the files at the given paths and serves it to the client
func ServeZip(w http.ResponseWriter, files_paths []string, zip_file_name string) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+zip_file_name)
	// Create a zip archive
	zip_writer := zip.NewWriter(w)
	defer func(zip_writer *zip.Writer) {
		err := zip_writer.Close()
		if err != nil {
			logger.Error("Error closing zip writer: ", err)
		}
	}(zip_writer)
	for _, file_path := range files_paths {
		err := addFileToZipWriter(zip_writer, file_path)
		if err != nil {
			SendErrorToClient(w, NewInternalServerError("Error adding file to zip: "+err.Error()))
			return
		}
	}
}

// ServeSubdirsZip creates a zip file from all the files given and organizes them in subdirectories
func ServeSubdirsZip(w http.ResponseWriter, files_paths map[string][]string, zip_file_name string) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+zip_file_name)
	// Create a zip archive
	zip_writer := zip.NewWriter(w)
	defer func(zip_writer *zip.Writer) {
		err := zip_writer.Close()
		if err != nil {
			logger.Error("Error closing zip writer: ", err)
		}
	}(zip_writer)
	for subdir, files := range files_paths {
		subdir_writer, err := zip_writer.Create(subdir)
		if err != nil {
			SendErrorToClient(w, NewInternalServerError("Error creating subdir in zip: "+err.Error()))
			return
		}
		subdir_zip_writer := zip.NewWriter(subdir_writer)
		for _, file_path := range files {
			err := addFileToZipWriter(subdir_zip_writer, file_path)
			if err != nil {
				SendErrorToClient(w, NewInternalServerError("Error adding file to zip: "+err.Error()))
				return
			}
		}
		err = subdir_zip_writer.Close()
		if err != nil {
			SendErrorToClient(w, NewInternalServerError("Error closing subdir zip writer: "+err.Error()))
			return
		}
	}
}

// addFileToZipWriter adds a file to a zip.Writer
func addFileToZipWriter(zip_writer *zip.Writer, file_path string) error {
	file, err := zip_writer.Create(file_path)
	if err != nil {
		return err
	}
	// Open the file
	file_to_zip, err := os.Open(file_path)
	if err != nil {
		return err
	}
	defer func(file_to_zip *os.File) {
		err := file_to_zip.Close()
		if err != nil {
			logger.Error("Error closing file: ", err)
		}
	}(file_to_zip)
	_, err = file_to_zip.WriteTo(file)
	if err != nil {
		return err
	}
	return nil
}

func SetCookie(w http.ResponseWriter, name string, value string, path string, secure bool, http_only bool, expires time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Secure:   secure,
		HttpOnly: http_only,
		Expires:  expires,
	})
}

func SetAuthCookie(w http.ResponseWriter, user_id int, auth_token string) {
	jwt, err := auth_utils.CreateAuthJWT(user_id, auth_token)
	if err != nil {
		logger.Error("Error creating auth JWT: ", err)
		return
	}
	SetCookie(w, "session_token", jwt, "/", true, true, time.Now().Add(auth_utils.DEFAULT_TOKEN_EXPIRATION))
}
