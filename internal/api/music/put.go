/*
Contains the handler for the PUT requests to the /musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	"encoding/json"
	"net/http"
)

// putMusicsHandler is the handler for the PUT /musics endpoint
// It updates the fields of a music in the database
func putMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a PUT request
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the music ID from the URL
	music_id := r.URL.Path[len("/musics/"):]

	// Parse the JSON body of the request
	decoder := json.NewDecoder(r.Body)
	var music map[string]interface{}
	err := decoder.Decode(&music)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Update the music in the database
	err = music_controller.UpdateMusic(music_id, music)
	if err != nil {
		http.Error(w, "Internal server error when updating music", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
