/*
Contains the handler for the POST requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	"net/http"
)

const MAX_UPLOAD_SIZE = 25 * 1024 * 1024 // 25Mb

// POST music handler
// Checks that the request is under 20Mb and that the file is a valid .mp3 file
// Saves the file to the server
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data to check if the request total size is under 20Mb
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE + 512)
	if err != nil {
		http.Error(w, "File too big (Max 20Mb)", http.StatusRequestEntityTooLarge)
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
		music_controller.PostMusic(title, author, genres, album, file)
	}
}
