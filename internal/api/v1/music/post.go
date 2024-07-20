/*
Contains the handler for the POST requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"
)

func AUTHORIZED_GENRES() []string {
	return []string{"pop", "rock", "rap", "jazz", "classical", "electro", "reggae", "metal", "blues", "country", "folk", "disco", "funk", "hip-hop", "house", "techno", "soul", "rnb", "reggaeton", "dancehall", "trap", "dubstep", "drumnbass", "hardstyle", "ambient", "chill", "dub", "garage", "grime", "indie", "jungle", "k-pop", "latin", "newage", "opera", "punk", "ska", "trance", "world", "other"}
}

func checkMusicGenres(genres []string) bool {
	for _, genre := range genres {
		if !format_utils.CheckStringInArray(AUTHORIZED_GENRES(), genre) {
			return false
		}
	}
	return true
}

// uploadHandler creates a new music
// @ Summary: Creates a new music
// @ Description: Creates a new music
// @ Tags: musics
// @ Accept: json
// @ Produces: json
// @ Param: title formData string true "Music title"
// @ Param: artists_ids formData []int true "Artist ID"
// @ Param: genres formData []string true "Music genres"
// @ Param: album_id formData int false "Album ID"
// @ Param: music formData file true "Music file"
// @ Param: illustration formData file false "Music illustration"
// @ Failure 400 {string} string "No music file found for key 'music'"
// @ Failure 400 {string} string "Invalid music file (must be audio/mpeg)"
// @ Failure 400 {string} string "Invalid artists_ids (must be positive integers)"
// @ Failure 400 {string} string "Missing at least 1 of required fields (title & author(s))"
// @ Failure 400 {string} string "Invalid album_id (must be > -1 integer)"
// @ Success 201 {string} string "Created"
// @ Router /api/musics [post]
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	if err != nil {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}
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

	// Check if illustration file is valid
	illustration_file_name := file_utils.DEFAULT_ILLUSTRATION_FILE
	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		illustration_file_name, _ = file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "albums")
	}

	// Check if other fields are valid
	title := r.FormValue("title")
	artists_ids_str := r.Form["artists_ids"]
	artists_ids, err := format_utils.ConvertStringArrayToIntArray(artists_ids_str)
	if err != nil {
		http.Error(w, "Invalid artists_ids (must be positive integers)", http.StatusBadRequest)
		return
	}
	if title == "" || len(artists_ids) < 1 {
		http.Error(w, "Missing at least 1 of required fields (title & author(s))", http.StatusBadRequest)
		return
	}

	genres := r.Form["genres"]
	if !checkMusicGenres(genres) {
		http.Error(w, "Invalid genres selected", http.StatusBadRequest)
		return
	}

	album_id_str := r.FormValue("album_id")
	album_id := -1
	if album_id_str != "" {
		album_id, err = strconv.Atoi(album_id_str)
		if err != nil || album_id < 0 {
			http.Error(w, "Invalid album_id (must be > -1 integer)", http.StatusBadRequest)
			return
		}
	}

	music_id, err := music_controller.PostMusic(title, genres, album_id, music_file, illustration_file_name)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.WriteHeader(http.StatusCreated)

	go music_controller.AddArtistsToMusic(music_id, artists_ids) // Don't wait for the artists to be added
}
