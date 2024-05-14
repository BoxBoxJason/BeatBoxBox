package artist_handler

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getArtistsHandler returns a list of artists
// @Summary Returns a list of artists
// @Description Returns a list of artists
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artists_ids query string true "Artists ids"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/artists [get]
func getArtistsHandler(w http.ResponseWriter, r *http.Request) {
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

	artists, err := artist_controller.GetArtists(artists_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(artists)
}

// getArtistHandler returns an artist given an id
// @Summary Returns an artist given an id
// @Description Returns an artist given an id
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artist_id path int true "Artist ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/artists/{artist_id} [get]
func getArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist_id_str := mux.Vars(r)["artist_id"]
	artist_id, err := strconv.Atoi(artist_id_str)
	if err != nil {
		http.Error(w, "artist id must be a positive integer", http.StatusBadRequest)
		return
	}

	artist, err := artist_controller.GetArtist(artist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(artist)
}
