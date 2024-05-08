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

func addMusicToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	err = playlist_controller.AddMusicsToPlaylist(playlist_id, []int{music_id})
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func addMusicsToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	music_ids_str := r.URL.Query().Get("music_ids")
	music_ids := []int{}
	if music_ids_str != "" {
		music_ids_str := strings.Split(music_ids_str, ",")
		for _, music_id := range music_ids_str {
			music_id, err := strconv.Atoi(music_id)
			if err != nil {
				http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
				return
			}
			music_ids = append(music_ids, music_id)
		}
	}

	err = playlist_controller.AddMusicsToPlaylist(playlist_id, music_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func removeMusicFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	err = playlist_controller.RemoveMusicsFromPlaylist(playlist_id, []int{music_id})
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func removeMusicsFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	music_ids_str := r.URL.Query().Get("music_ids")
	music_ids := []int{}
	if music_ids_str != "" {
		music_ids_str := strings.Split(music_ids_str, ",")
		for _, music_id := range music_ids_str {
			music_id, err := strconv.Atoi(music_id)
			if err != nil {
				http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
				return
			}
			music_ids = append(music_ids, music_id)
		}
	}

	err = playlist_controller.RemoveMusicsFromPlaylist(playlist_id, music_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func putPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	update_dict, err := parseURLParams(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	illustration_file, illustration_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		new_illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_header, illustration_file, "playlists")
		if err != nil || new_illustration_file_name == "" {
			custom_errors.SendErrorToClient(err, w, "")
			return
		}
		update_dict["illustration"] = new_illustration_file_name
	}
	err = playlist_controller.UpdatePlaylist(playlist_id, update_dict)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
