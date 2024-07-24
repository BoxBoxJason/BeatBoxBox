package playlist_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// addMusicToPlaylistHandler adds a music to the playlist with the given ID
// @Summary Add a music to a playlist by their IDs
// @Description Add a music to a playlist by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Param music_id path int true "Music Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when adding music to playlist"
// @Router /api/playlists/{playlist_id}/add/{music_id} [put]
func addMusicToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
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

	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicExists(music_id) {
		http.Error(w, "Music does not exist", http.StatusNotFound)
		return
	}

	err = playlist_controller.AddMusicsToPlaylist(playlist_id, []int{music_id})
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// addMusicsToPlaylistHandler adds musics to the playlist with the given ID
// @Summary Add musics to a playlist by their IDs
// @Description Add musics to a playlist by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Param music_ids query string true "Music Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when adding musics to playlist"
// @Router /api/playlists/{playlist_id}/add [put]
func addMusicsToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
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

	music_ids_str := r.URL.Query().Get("music_ids")
	music_ids, err := format_utils.ConvertStringToIntArray(music_ids_str, ",")
	if err != nil {
		http.Error(w, "Invalid music IDs provided, please use a valid integer music IDs", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicsExists(music_ids) {
		http.Error(w, "At least one music does not exist", http.StatusNotFound)
		return
	}

	err = playlist_controller.AddMusicsToPlaylist(playlist_id, music_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// removeMusicFromPlaylistHandler removes a music from the playlist with the given ID
// @Summary Remove a music from a playlist by their IDs
// @Description Remove a music from a playlist by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Param music_id path int true "Music Id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when removing music from playlist"
// @Router /api/playlists/{playlist_id}/remove/{music_id} [put]
func removeMusicFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
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

	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	err = playlist_controller.RemoveMusicsFromPlaylist(playlist_id, []int{music_id})
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// removeMusicsFromPlaylistHandler removes musics from the playlist with the given ID
// @Summary Remove musics from a playlist by their IDs
// @Description Remove musics from a playlist by their IDs
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Param music_ids query string true "Music Ids"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 404 {string} string "Playlist does not exist"
// @Failure 500 {string} string "Internal server error when removing musics from playlist"
// @Router /api/playlists/{playlist_id}/remove [put]
func removeMusicsFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
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

	music_ids_str := r.URL.Query().Get("music_ids")
	music_ids, err := format_utils.ConvertStringToIntArray(music_ids_str, ",")
	if err != nil {
		http.Error(w, "Invalid music IDs provided, please use a valid integer music IDs", http.StatusBadRequest)
		return
	}

	err = playlist_controller.RemoveMusicsFromPlaylist(playlist_id, music_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// putPlaylistHandler updates the playlist with the given ID
// @Summary Update a playlist by its ID
// @Description Update a playlist by its ID
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist_id path int true "Playlist Id"
// @Param title formData string false "Title"
// @Param description formData string false "Description"
// @Param illustration formData file false "Illustration"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid playlist ID provided, please use a valid integer playlist ID"
// @Failure 500 {string} string "Internal server error when updating playlist"
// @Router /api/playlists/{playlist_id} [put]
func putPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	playlist_id_str := mux.Vars(r)["playlist_id"]
	playlist_id, err := strconv.Atoi(playlist_id_str)
	if err != nil {
		http.Error(w, "Invalid playlist ID provided, please use a valid integer playlist ID", http.StatusBadRequest)
		return
	}

	update_dict := make(map[string]interface{})
	title := r.FormValue("title")
	if title != "" {
		update_dict["title"] = title
	}
	description := r.FormValue("description")
	if description != "" {
		update_dict["description"] = description
	}

	illustration_file, illustration_header, err := r.FormFile("illustration")
	if err == nil {
		defer illustration_file.Close()
		new_illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_header, illustration_file, "playlists")
		if err != nil || new_illustration_file_name == "" {
			custom_errors.SendErrorToClient(err, w, "")
			return
		}
		update_dict["illustration"] = new_illustration_file_name
	}
	err = playlist_controller.UpdatePlaylist(playlist_id, update_dict)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
