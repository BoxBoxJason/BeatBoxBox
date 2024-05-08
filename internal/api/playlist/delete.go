package playlist_handler

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	err = playlist_controller.DeletePlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deletePlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	raw_playlists_ids := r.URL.Query().Get("playlists_ids")
	playlists_ids := []int{}

	if raw_playlists_ids != "" {
		playlists_ids_str := strings.Split(raw_playlists_ids, ",")
		for _, playlist_id := range playlists_ids_str {
			playlist_id, err := strconv.Atoi(playlist_id)
			if err != nil {
				http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
				return
			}
			playlists_ids = append(playlists_ids, playlist_id)
		}

		err := playlist_controller.DeletePlaylists(playlists_ids)
		if err != nil {
			custom_errors.SendErrorToClient(err, w, "")
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
