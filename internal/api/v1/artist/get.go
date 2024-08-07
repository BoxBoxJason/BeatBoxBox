package artist_handler_v1

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/utils/httputils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist_id, err := strconv.Atoi(mux.Vars(r)["artist_id"])
	if err != nil {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid artist ID"))
		return
	} else if artist_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid artist ID: %d", artist_id)))
		return
	}
	artist, err := artist_controller.GetArtistJSON(artist_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, artist)
}

func getArtistsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{"pseudo", "partial_pseudo", "genre", "album", "music"}, []string{"artist_id", "album_id", "music_id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	var artists []byte
	if params["artist_id"] != nil && len(params) > 1 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Can't use artist_id with other parameters"))
		return
	} else if params["artist_id"] != nil {
		artists_ids := params["artist_id"].([]int)
		artists, err = artist_controller.GetArtistsJSON(artists_ids)
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
	} else {
		pseudos := params["pseudo"].([]string)
		partial_pseudos := params["partial_pseudo"].([]string)
		genres := params["genre"].([]string)
		albums_ids := params["album_id"].([]int)
		albums_titles := params["album"].([]string)
		musics_ids := params["music_id"].([]int)
		musics_titles := params["music"].([]string)
		artists, err = artist_controller.GetArtistsJSONFromFilters(pseudos, partial_pseudos, genres, albums_ids, albums_titles, musics_ids, musics_titles)
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
	}
	httputils.RespondWithJSON(w, http.StatusOK, artists)
}
