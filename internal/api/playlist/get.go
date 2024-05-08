package playlist_handler

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func downloadPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	title, musics_paths, err := playlist_controller.GetMusicsPathFromPlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServeZip(w, musics_paths, title+".zip")
}

func downloadPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	playlist_ids_str := r.URL.Query().Get("playlist_ids")
	playlist_ids := []int{}
	if playlist_ids_str != "" {
		playlist_ids_str := strings.Split(playlist_ids_str, ",")
		for _, playlist_id_str := range playlist_ids_str {
			playlist_id, err := strconv.Atoi(playlist_id_str)
			if err != nil {
				http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
				return
			}
			playlist_ids = append(playlist_ids, playlist_id)
		}
	}

	musics_paths, err := playlist_controller.GetMusicsPathFromPlaylists(playlist_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServePlaylistsZip(w, musics_paths)
}

func getPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	playlist_json, err := playlist_controller.GetPlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(playlist_json)
	w.WriteHeader(http.StatusOK)
}

func getPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	playlist_ids_str := r.URL.Query().Get("playlist_ids")
	playlist_ids := []int{}
	if playlist_ids_str != "" {
		playlist_ids_str := strings.Split(playlist_ids_str, ",")
		for _, playlist_id_str := range playlist_ids_str {
			playlist_id, err := strconv.Atoi(playlist_id_str)
			if err != nil {
				http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
				return
			}
			playlist_ids = append(playlist_ids, playlist_id)
		}
	}

	playlists_json, err := playlist_controller.GetPlaylists(playlist_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(playlists_json)
	w.WriteHeader(http.StatusOK)
}
