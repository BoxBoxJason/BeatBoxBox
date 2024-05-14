package playlist_handler

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getPlaylistHandler returns the playlist with the given ID
// @Summary Get a playlist by its ID
// @Description Get a playlist by its ID
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when getting playlist"
// @Router /api/playlists/{playlist_id} [get]
func downloadPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}
	if !playlist_controller.PlaylistExists(playlist_id) {
		http.Error(w, "Playlist does not exist", http.StatusNotFound)
		return
	}

	title, musics_paths, err := playlist_controller.GetMusicsPathFromPlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServeZip(w, musics_paths, title+".zip")
}

// getPlaylistsHandler returns the playlists with the given IDs
// @Summary Get playlists by their IDs
// @Description Get playlists by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_ids query string true "Playlist Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist IDs provided, please use a valid integer playlist IDs"
// @Failure 404 {string} string "At least one playlist does not exist"
// @Failure 500 {string} string "Internal server error when getting playlists"
// @Router /api/playlists [get]
func downloadPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	playlist_ids_str := r.URL.Query().Get("playlist_ids")
	playlist_ids, err := format_utils.ConvertStringToIntArray(playlist_ids_str, ",")
	if err != nil {
		http.Error(w, "Invalid playlist IDs provided, please use a valid integer playlist IDs", http.StatusBadRequest)
		return
	}
	if !playlist_controller.PlaylistsExist(playlist_ids) {
		http.Error(w, "At least one playlist does not exist", http.StatusNotFound)
		return
	}

	musics_paths, err := playlist_controller.GetMusicsPathFromPlaylists(playlist_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServeTreeZip(w, musics_paths, "playlists")
}

// getPlaylistHandler returns the playlist with the given ID
// @Summary Get a playlist by its ID
// @Description Get a playlist by its ID
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when getting playlist"
// @Router /api/playlists/{playlist_id} [get]
func getPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}
	if !playlist_controller.PlaylistExists(playlist_id) {
		http.Error(w, "Playlist does not exist", http.StatusNotFound)
		return
	}

	playlist_json, err := playlist_controller.GetPlaylist(playlist_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(playlist_json)
}

// getPlaylistsHandler returns the playlists with the given IDs
// @Summary Get playlists by their IDs
// @Description Get playlists by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_ids query string true "Playlist Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist IDs provided, please use a valid integer playlist IDs"
// @Failure 404 {string} string "At least one playlist does not exist"
// @Failure 500 {string} string "Internal server error when getting playlists"
// @Router /api/playlists [get]
func getPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	playlist_ids_str := r.URL.Query().Get("playlist_ids")
	playlist_ids, err := format_utils.ConvertStringToIntArray(playlist_ids_str, ",")
	if err != nil {
		http.Error(w, "Invalid playlist IDs provided, please use a valid integer playlist IDs", http.StatusBadRequest)
		return
	}
	if !playlist_controller.PlaylistsExist(playlist_ids) {
		http.Error(w, "At least one playlist does not exist", http.StatusNotFound)
		return
	}

	playlists_json, err := playlist_controller.GetPlaylists(playlist_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(playlists_json)
}
