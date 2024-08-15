package artist_handler_v1

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	vars := mux.Vars(r)
	artist_id, err := strconv.Atoi(vars["artist_id"])
	if err != nil || artist_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("artist_id is required and must be a positive integer"))
		return
	}
	// Delete the artist
	err = artist_controller.DeleteArtist(artist_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	artists_ids := params["id"].([]int)
	// Delete the artists
	err = artist_controller.DeleteArtists(artists_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
