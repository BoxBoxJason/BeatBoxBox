package artist_handler

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
)

// postArtistHandler creates a new artist
// @Summary Creates a new artist
// @Description Creates a new artist
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artist_name formData string true "Artist name"
// @Param illustration formData file false "Illustration"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /api/artists [post]
func postArtistHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	if err != nil {
		http.Error(w, "Request too large", http.StatusBadRequest)
		return
	}

	artist_name := r.FormValue("artist_name")
	if artist_name == "" {
		http.Error(w, "Artist name is required", http.StatusBadRequest)
		return
	}

	illustration_file_name := file_utils.DEFAULT_ILLUSTRATION_FILE
	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		illustration_file_name, _ = file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "artists")
	}
	_, err = artist_controller.PostArtist(artist_name, illustration_file_name)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.WriteHeader(http.StatusCreated)
}
