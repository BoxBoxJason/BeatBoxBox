package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAlbumHandler(w http.ResponseWriter, r *http.Request) {
	album_id, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil || album_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid album_id, it must be a positive integer"))
		return
	}
	album, err := album_controller.GetAlbumJSON(album_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, album)
}

func getAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{"title", "partial_title", "artist", "music", "genre"}, []string{"album_id", "artist_id", "music_id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	var albums []byte
	// Filter request if incompatible parameters are used
	if params["album_id"] != nil && len(params) > 1 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Can't use album_id with other parameters"))
		return
	} else if params["album_id"] != nil {
		albums_ids := params["album_id"].([]int)
		if len(albums_ids) > 0 {
			albums, err = album_controller.GetAlbumsJSON(albums_ids)
			if err != nil {
				custom_errors.SendErrorToClient(w, err)
				return
			}
		}
	} else {
		titles := params["title"].([]string)
		partial_titles := params["partial_title"].([]string)
		genres := params["genre"].([]string)
		artists_names := params["artist"].([]string)
		musics_names := params["music"].([]string)
		artists_ids := params["artist_id"].([]int)
		musics_ids := params["music_id"].([]int)
		albums, err = album_controller.GetAlbumsJSONFromFilters(titles, partial_titles, genres, artists_names, musics_names, artists_ids, musics_ids)
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
	}
	httputils.RespondWithJSON(w, http.StatusOK, albums)
}
