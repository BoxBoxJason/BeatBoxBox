/*
Contains the handler for the GET requests to the /api/musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// getMusicHandler returns the music with the given ID
// @Summary Get a music by its ID
// @Description Get a music by its ID
// @Tags musics
// @Accept json
// @Produce json
// @Param music_id path int true "Music Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @Failure 404 {string} string "Music does not exist"
// @Failure 500 {string} string "Unexpected error while getting music"
// @Router /api/musics/{music_id} [get]
func getMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicExists(music_id) {
		http.Error(w, "Music does not exist", http.StatusNotFound)
		return
	}

	// Get the music from the database
	music, err := music_controller.GetMusic(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the music as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(music)
	w.WriteHeader(http.StatusOK)
}

// getMusicsHandler returns the musics with the given IDs
// @Summary Get musics by their IDs
// @Description Get musics by their IDs
// @Tags musics
// @Accept json
// @Produce json
// @Param music_ids query []int true "Musics Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "music_ids must be comma separated integers"
// @Failure 404 {string} string "One or more musics do not exist"
// @Failure 500 {string} string "Unexpected error while getting musics"
// @Router /api/musics [get]
func getMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve requested music IDs from the URL
	musics_ids_str := r.URL.Query().Get("music_ids")
	if musics_ids_str == "" {
		http.Error(w, "No music ID provided, please use music_ids request parameter", http.StatusBadRequest)
		return
	}

	musics_ids, err := format_utils.ConvertStringToIntArray(musics_ids_str, ",")
	if err != nil {
		http.Error(w, "music_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicsExists(musics_ids) {
		http.Error(w, "One or more musics do not exist", http.StatusNotFound)
		return
	}

	// Get the musics from the database
	musics_json, err := music_controller.GetMusics(musics_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the musics as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(musics_json)
	w.WriteHeader(http.StatusOK)
}

// downloadMusicHandler returns the music file with the given ID
// @Summary Download a music by its ID
// @Description Download a music by its ID
// @Tags musics
// @Accept json
// @Produce json
// @Param music_id path int true "Music Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @Failure 404 {string} string "Music does not exist"
// @Failure 500 {string} string "Unexpected error while downloading music"
// @Router /api/musics/{music_id}/download [get]
func downloadMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL & retrieve the corresponding music path
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicExists(music_id) {
		http.Error(w, "Music does not exist", http.StatusNotFound)
		return
	}

	music_path, err := music_controller.GetMusicPathFromId(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Open the music file
	music_file, err := os.Open(music_path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer music_file.Close()

	// Serve the file
	http.ServeFile(w, r, music_path)
	w.WriteHeader(http.StatusOK)
}

// downloadMusicsHandler downloads the musics with the given IDs
// @Summary Download musics by their IDs
// @Description Download musics by their IDs
// @Tags musics
// @Accept json
// @Produce json
// @Param music_ids query []int true "Musics Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "music_ids must be comma separated integers"
// @Failure 404 {string} string "One or more musics do not exist"
// @Failure 500 {string} string "Unexpected error while downloading musics"
// @Router /api/musics/download [get]
func downloadMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL & retrieve the corresponding music path
	music_ids_requested := r.URL.Query().Get("music_ids")
	if music_ids_requested == "" {
		http.Error(w, "No music ID provided, please use music_ids request parameter", http.StatusBadRequest)
		return
	}

	music_ids_str := strings.Split(music_ids_requested, ",")
	if len(music_ids_str) == 0 {
		http.Error(w, "No music ID provided, please use music_ids request parameter", http.StatusBadRequest)
		return
	}
	musics_ids, err := format_utils.ConvertStringToIntArray(music_ids_requested, ",")
	if err != nil {
		http.Error(w, "music_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicsExists(musics_ids) {
		http.Error(w, "One or more musics do not exist", http.StatusNotFound)
		return
	}

	musics_paths, err := music_controller.GetMusicsPathFromIds(musics_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file_utils.ServeZip(w, musics_paths, "musics.zip")
}
