package playlist_handler

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
	"strconv"
)

func createPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	user_id, _ := auth_middleware.AuthenticateUser(r)
	if user_id < 0 {
		http.Redirect(w, r, "/auth", http.StatusUnauthorized)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "title is empty", http.StatusBadRequest)
		return
	}
	description := r.FormValue("description")

	// Check if illustration file is valid
	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		err = file_utils.CheckFileMeetsRequirements(*illustration_file_header, file_utils.MAX_IMAGE_FILE_SIZE, "image/jpeg")
		if err != nil {
			custom_errors.SendErrorToClient(err, w, "")
			return
		}
	} else {
		illustration_file = nil
	}

	playlist_id, err := playlist_controller.PostPlaylist(title, user_id, description, illustration_file)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	musics_ids_str := r.Form["musics_ids"]
	if len(musics_ids_str) > 0 {
		musics_ids := []int{}
		for _, music_id_str := range musics_ids_str {
			music_id, err := strconv.Atoi(music_id_str)
			if err != nil {
				http.Error(w, "music_id is not a number", http.StatusBadRequest)
				return
			}
			musics_ids = append(musics_ids, music_id)
		}
		go playlist_controller.AddMusicsToPlaylist(playlist_id, musics_ids) // Do not wait for the musics to be added to the album
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
