package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil || playlist_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid playlist ID, must be a positive integer"))
		return
	}
	playlist_json, err := playlist_controller.GetPlaylistJSON(playlist_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, playlist_json)
}

func downloadPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Get the playlist ID from the URL
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil || playlist_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid playlist ID, must be a positive integer"))
		return
	}
	err = playlist_controller.ServePlaylistFiles(w, playlist_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
}

func getPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, nil, nil, []string{"title", "music", "owner"}, []string{"id", "music_id", "owner_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	var playlists []byte
	if params["id"] != nil && len(params) > 1 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid query parameters: id cannot be used with other parameters"))
		return
	} else if params["id"] != nil {
		playlists, err = playlist_controller.GetPlaylistsJSON(params["id"].([]int))
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	} else {
		titles := params["title"].([]string)
		musics := params["music"].([]string)
		owners := params["owner"].([]string)
		music_ids := params["music_id"].([]int)
		owner_ids := params["owner_id"].([]int)
		playlists, err = playlist_controller.GetPlaylistsJSONByFilters(titles, musics, owners, music_ids, owner_ids)
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	}
	httputils.RespondWithJSON(w, http.StatusOK, playlists)

}

func downloadPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if len(params) == 0 || len(params["id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("No playlist ID provided"))
		return
	}
	playlist_ids := params["id"].([]int)
	err = playlist_controller.ServePlaylistsFiles(w, playlist_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
}
