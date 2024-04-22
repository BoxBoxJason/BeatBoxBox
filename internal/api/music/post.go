/*
Contains the handler for the POST requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	"net/http"
	"strconv"
)

const MAX_UPLOAD_SIZE = 25 * 1024 * 1024 // 25Mb

// POST music handler
// Checks that the request is under 20Mb and that the file is a valid .mp3 file
// Saves the file to the server
func uploadHandler(w http.ResponseWriter, r *http.Request) {
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
	artist_id_str := r.FormValue("artist_id")
	genres := r.Form["genres"]
	album_id_str := r.FormValue("album_id")

	if title == "" || artist_id_str == "" {
		http.Error(w, "Missing required fields (title & author)", http.StatusBadRequest)
		return
	}

	// Check that the artist_id and album_id are valid
	artist_id, err := strconv.Atoi(artist_id_str)
	if err != nil || artist_id < 0 {
		http.Error(w, "Invalid artist_id (must be positive integer)", http.StatusBadRequest)
		return
	}

	album_id := -1
	if album_id_str != "" {
		album_id, err = strconv.Atoi(album_id_str)
		if err != nil || album_id < -1 {
			http.Error(w, "Invalid album_id (must be >= -1 integer)", http.StatusBadRequest)
			return
		}
	}

	// Check if the file is a valid .mp3 file
	if file_header.Header.Get("Content-Type") != "audio/mpeg" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	} else {
		// Save the file to the server
		music_controller.PostMusic(title, artist_id, genres, album_id, file)
	}
}
