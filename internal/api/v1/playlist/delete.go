package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil || playlist_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid playlist ID, must be a positive integer"))
		return
	}
	err = playlist_controller.DeletePlaylist(playlist_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deletePlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if len(params) == 0 || len(params["id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("No playlist ID provided"))
		return
	}
	playlist_ids := params["id"].([]int)
	err = playlist_controller.DeletePlaylists(playlist_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	w.WriteHeader(http.StatusNoContent)
}
