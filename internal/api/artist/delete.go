package artist_handler

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// deleteArtistsHandler deletes artist with the given id
// @Summary Deletes artist with the given id
// @Description Deletes artist with the given id
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artist_id path int true "Artist ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/artists/{artist_id} [delete]
func deleteArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist_id_str := mux.Vars(r)["artist_id"]
	artist_id, err := strconv.Atoi(artist_id_str)
	if err != nil {
		http.Error(w, "artist id must be a positive integer", http.StatusBadRequest)
		return
	}

	err = artist_controller.DeleteArtist(artist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteArtistsHandler deletes artists with the given ids
// @Summary Deletes artists with the given ids
// @Description Deletes artists with the given ids
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artists_ids query string true "Artists IDs"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/artists [delete]
func deleteArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists_ids_str := r.URL.Query().Get("artists_ids")
	if artists_ids_str == "" {
		http.Error(w, "artists_ids must be a list of positive integers", http.StatusBadRequest)
		return
	}
	artists_ids, err := format_utils.ConvertStringToIntArray(artists_ids_str, ",")
	if err != nil {
		http.Error(w, "artists_ids must be a list of positive integers", http.StatusBadRequest)
		return
	}

	err = artist_controller.DeleteArtists(artists_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
