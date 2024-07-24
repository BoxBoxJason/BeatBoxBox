package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// deletePlaylistHandler deletes the playlist with the given ID
// @Summary Delete a playlist by its ID
// @Description Delete a playlist by its ID
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when deleting playlist"
// @Router /api/playlists/{playlist_id} [delete]
func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	err = playlist_controller.DeletePlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deletePlaylistsHandler deletes the playlists with the given IDs
// @Summary Delete playlists by their IDs
// @Description Delete playlists by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlists_ids query string true "Playlist Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist IDs provided, please use a valid integer playlist IDs"
// @Failure 404 {string} string "At least one playlist does not exist"
// @Failure 500 {string} string "Internal server error when deleting playlists"
// @Router /api/playlists [delete]
func deletePlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	raw_playlists_ids := r.URL.Query().Get("playlists_ids")
	playlists_ids, err := format_utils.ConvertStringToIntArray(raw_playlists_ids, ",")
	if err != nil {
		http.Error(w, "Invalid playlist IDs provided, please use a valid integer playlist IDs", http.StatusBadRequest)
		return
	}
	if !playlist_controller.PlaylistsExist(playlists_ids) {
		http.Error(w, "At least one playlist does not exist", http.StatusNotFound)
		return
	}

	err = playlist_controller.DeletePlaylists(playlists_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.WriteHeader(http.StatusOK)
}
