package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
)

// createPlaylistHandler creates a new playlist
// @Summary Create a new playlist
// @Description Create a new playlist
// @Tags playlists
// @Accept json
// @Produce json
// @Param title formData string true "Title"
// @Param description formData string false "Description"
// @Param illustration formData file false "Illustration"
// @Param musics_ids formData []int false "Musics Ids"
// @Success 201 {string} string "Created"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "title is empty"
// @Failure 401 {string} string "Unauthorized"
// @Failure 413 {string} string "Request too large"
// @Router /api/playlists [post]
func createPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	if err != nil {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}
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
	illustration_file_name := file_utils.DEFAULT_ILLUSTRATION_FILE
	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		illustration_file_name, _ = file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "playlists")
	}

	playlist_id, err := playlist_controller.PostPlaylist(title, user_id, description, illustration_file_name)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	musics_ids_str := r.Form["musics_ids"]
	musics_ids, err := format_utils.ConvertStringArrayToIntArray(musics_ids_str)
	if err != nil {
		http.Error(w, "musics_ids must be integers", http.StatusBadRequest)
		return
	}
	if len(musics_ids) > 0 {
		go playlist_controller.AddMusicsToPlaylist(playlist_id, musics_ids) // Do not wait for the musics to be added to the album
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
