/*
Package music_handler is the handler for music API

This package is responsible for handling all the API requests related to music.
It creates a new router and registers all the handlers for the music API.
*/

package music_handler

import (
	"net/http"
)

const MAX_UPLOAD_SIZE = 20 * 1024 * 1024 // 20Mb
const MUSIC_DIRECTORY = "./data/musics"

// POST music handler, checks if the request is a POST request and then
// Checks that the request is under 20Mb and that the file is a valid .mp3 file
// Saves the file to the server
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data to check if the request total size is under 20Mb
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE + 512)
	if err != nil {
		http.Error(w, "File too big", http.StatusRequestEntityTooLarge)
		return
	}

	// Check that all the required fields are present
	file, file_header, err := r.FormFile("music")
	if err != nil {
		http.Error(w, "No file found", http.StatusBadRequest)
		return
	}
	defer file.Close()
	title := r.FormValue("title")
	author := r.FormValue("author")
	genres := r.Form["genres"]
	album := r.FormValue("album")

	if title == "" || author == "" {
		http.Error(w, "Missing required fields (title & author)", http.StatusBadRequest)
		return
	}
	// Check if the file is a valid .mp3 file
	if file_header.Header.Get("Content-Type") != "audio/mpeg" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	} else {
		// Save the file to the server

	}

}

// GET musics handler, checks if the request is a GET request and then
// Returns all the musics in the database as a JSON response
// Parameters can be passed to filter the musics
func getMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get parameters from the request
	params := r.URL.Query()

	// Get the musics from the database

	// Return the musics as a JSON response
}
