package music_handler

import (
	"net/http"
	"strconv"
	"strings"

	music_controller "BeatBoxBox/internal/controller/music"

	"github.com/gorilla/mux"
)

func deleteMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	// Delete the music from the database
	err = music_controller.DeleteMusic(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE musics handler,
// Deletes selected musics from the database based on the music_ids request parameter
func deleteMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the music IDs from the request url
	raw_music_ids_str := r.URL.Query().Get("music_ids")
	music_ids := []int{}

	if raw_music_ids_str != "" {
		music_ids_str := strings.Split(raw_music_ids_str, ",")
		for _, music_id_str := range music_ids_str {
			music_id, err := strconv.Atoi(music_id_str)
			if err != nil {
				http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
				return
			}
			music_ids = append(music_ids, music_id)
		}
	}

	// Delete the musics from the database
	err := music_controller.DeleteMusics(music_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
