package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	artist_controller "BeatBoxBox/internal/controller/artist"
	music_controller "BeatBoxBox/internal/controller/music"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
)

// createAlbumHandler creates a new album
// @ Summary: Creates a new album
// @ Description: Creates a new album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: title formData string true "Album title"
// @ Param: description formData string false "Album description"
// @ Param: illustration formData file false "Album illustration"
// @ Param: artists_ids formData []int true "Artist ID"
// @ Param: musics_ids formData []int false "Music ID"
// @ Failure 400 {string} string "title is missing"
// @ Failure 400 {string} string "artists_ids must be integers"
// @ Failure 400 {string} string "At least one artist does not exist"
// @ Failure 400 {string} string "musics_ids must be integers"
// @ Failure 400 {string} string "At least one music does not exist"
// @ Failure 401 {string} string "Unauthorized"
// @ Success 201 {string} string "Created"
// @ Success 202 {string} string "Accepted"
// @ Router /api/albums [post]
func createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	if err != nil {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}

	// Check if user is authenticated
	user_id, _ := auth_middleware.AuthenticateUser(r)
	if user_id < 0 {
		http.Redirect(w, r, "/auth", http.StatusUnauthorized)
		return
	}

	// Check if title is valid
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "title is missing", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")

	// Check if artists_ids is valid
	artists_ids_str := r.Form["artists_ids"]
	artists_ids, err := format_utils.ConvertStringArrayToIntArray(artists_ids_str)
	if err != nil {
		http.Error(w, "artists_ids must be integers", http.StatusBadRequest)
		return
	}
	if !artist_controller.ArtistsExist(artists_ids) {
		http.Error(w, "At least one artist does not exist", http.StatusBadRequest)
		return
	}

	// Check if illustration file is valid
	illustration_file_name := file_utils.DEFAULT_ILLUSTRATION_FILE
	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		illustration_file_name, _ = file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "albums")
	}

	album_id, err := album_controller.PostAlbum(title, artists_ids, description, illustration_file_name)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.WriteHeader(http.StatusCreated)

	musics_ids_str := r.Form["musics_ids"]
	musics_ids, err := format_utils.ConvertStringArrayToIntArray(musics_ids_str)
	if err != nil {
		http.Error(w, "musics_ids must be integers", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicsExists(musics_ids) {
		http.Error(w, "At least one music does not exist", http.StatusBadRequest)
		return
	}
	go album_controller.AddMusicsToAlbum(album_id, musics_ids) // Do not wait for the musics to be added to the album
	w.WriteHeader(http.StatusAccepted)
}
