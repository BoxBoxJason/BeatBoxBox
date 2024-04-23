/*
Contains the handler for the PUT requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// putMusicsHandler is the handler for the PUT /musics endpoint
// It updates the fields of a music in the database
func putMusicsHandler(w http.ResponseWriter, r *http.Request) {

	// Get the music ID from the URL
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	// Parse the url parameters and retrieve only authorized ones
	update_dict, err := parseURLParams(r.URL.Query())
	if err != nil {
		http.Error(w, "Invalid URL parameters provided: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update the music in the database
	err = music_controller.UpdateMusic(music_id, update_dict)
	if err != nil {
		http.Error(w, "Internal server error when updating music: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
