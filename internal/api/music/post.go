/*
Contains the handler for the POST requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
	"strconv"
)

// POST music handler
// Checks that the request is under 20Mb and that the file is a valid .mp3 file
// Saves the file to the server
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check for file(s) size(s)
	music_file, music_file_header, err := r.FormFile("music")
	if err != nil {
		http.Error(w, "No music file found for key 'music'", http.StatusBadRequest)
		return
	}
	defer music_file.Close()
	err = file_utils.CheckFileMeetsRequirements(*music_file_header, file_utils.MAX_MUSIC_FILE_SIZE, "audio/mpeg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		err = file_utils.CheckFileMeetsRequirements(*illustration_file_header, file_utils.MAX_IMAGE_FILE_SIZE, "image/jpeg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		illustration_file = nil
	}

	// Check if other fields are valid
	title := r.FormValue("title")
	artist_id_str := r.FormValue("artist_id")
	genres := r.Form["genres"]
	album_id_str := r.FormValue("album_id")

	if title == "" || artist_id_str == "" {
		http.Error(w, "Missing at least 1 of required fields (title & author)", http.StatusBadRequest)
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
		if err != nil || album_id < 0 {
			http.Error(w, "Invalid album_id (must be > -1 integer)", http.StatusBadRequest)
			return
		}
	}

	_, err = music_controller.PostMusic(title, artist_id, genres, album_id, music_file, illustration_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
