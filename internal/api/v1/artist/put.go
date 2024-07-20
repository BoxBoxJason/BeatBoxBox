package artist_handler

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// putArtistHandler updates an artist with the given id
// @Summary Updates an artist with the given id
// @Description Updates an artist with the given id
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artist_id path int true "Artist ID"
// @Param artist_name formData string false "Artist Name"
// @Param illustration formData file false "Illustration"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/artists/{artist_id} [put]
func putArtistHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	artist_id_str := mux.Vars(r)["artist_id"]
	artist_id, err := strconv.Atoi(artist_id_str)
	if err != nil {
		http.Error(w, "artist id must be a positive integer", http.StatusBadRequest)
		return
	}

	update_dict := make(map[string]interface{})
	artist_name := r.FormValue("artist_name")
	if artist_name == "" {
		update_dict["pseudo"] = artist_name
	}

	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		illustration_file_name, _ := file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "artists")
		if illustration_file_name != file_utils.DEFAULT_ILLUSTRATION_FILE {
			update_dict["illustration"] = illustration_file_name
		}
	}

	err = artist_controller.UpdateArtist(artist_id, update_dict)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
