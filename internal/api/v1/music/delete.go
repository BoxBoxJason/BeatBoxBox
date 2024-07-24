package music_handler_v1

import (
	"net/http"
	"strconv"

	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"

	"github.com/gorilla/mux"
)

// deleteMusicHandler deletes the music with the given ID
// @Summary Delete a music by its ID
// @Description Delete a music by its ID
// @Tags musics
// @Accept json
// @Produce json
// @Param music_id path int true "Music Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @Failure 404 {string} string "Music does not exist"
// @Failure 500 {string} string "Unexpected error while deleting music"
// @Router /api/musics/{music_id} [delete]
func deleteMusicHandler(w http.ResponseWriter, r *http.Request) {
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

	// Delete the music from the database
	err = music_controller.DeleteMusic(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteMusicsHandler deletes the musics with the given IDs
// @Summary Delete musics by their IDs
// @Description Delete musics by their IDs
// @Tags musics
// @Accept json
// @Produce json
// @Param music_ids query []int true "Musics Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "music_ids must be comma separated integers"
// @Failure 404 {string} string "One or more musics do not exist"
// @Failure 500 {string} string "Unexpected error while deleting musics"
// @Router /api/musics [delete]
func deleteMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the music IDs from the request url
	raw_music_ids_str := r.URL.Query().Get("music_ids")
	musics_ids, err := format_utils.ConvertStringToIntArray(raw_music_ids_str, ",")
	if err != nil {
		http.Error(w, "music_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	if len(musics_ids) < 1 {
		http.Error(w, "At least one music id must be provided", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicsExists(musics_ids) {
		http.Error(w, "One or more musics do not exist", http.StatusNotFound)
		return
	}

	// Delete the musics from the database
	err = music_controller.DeleteMusics(musics_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
