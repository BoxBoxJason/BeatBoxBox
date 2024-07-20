package album_handler

import (
	album_controller "BeatBoxBox/internal/controller/album"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getAlbumHandler returns the album with the given ID
// @ Summary: Get an album by its ID
// @ Description: Get an album by its ID
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album Id"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 500 {string} string "Unexpected error while getting album"
// @ Router /api/albums/{album_id} [get]
func getAlbumHandler(w http.ResponseWriter, r *http.Request) {
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	album_json, err := album_controller.GetAlbum(album_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	w.Write(album_json)
}

// getAlbumsHandler returns the albums with the given IDs
// @ Summary: Get albums by their IDs
// @ Description: Get albums by their IDs
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: albums_ids query []int true "Albums Ids"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "albums_ids must be comma separated integers"
// @ Failure 500 {string} string "Unexpected error while getting albums"
// @ Router /api/albums [get]
func getAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albums_ids_str := r.URL.Query().Get("albums_ids")
	albums_ids, err := format_utils.ConvertStringToIntArray(albums_ids_str, ",")
	if err != nil {
		http.Error(w, "albums_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	albums_json, err := album_controller.GetAlbums(albums_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.Write(albums_json)
}

// downloadAlbumHandler downloads the album with the given ID
// @ Summary: Download an album by its ID
// @ Description: Download an album by its ID
// @ Tags: albums
// @ Accept: json
// @ Produces: zip
// @ Param: album_id path int true "Album Id"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 500 {string} string "Unexpected error while downloading album"
// @ Router /api/albums/{album_id}/download [get]
func downloadAlbumHandler(w http.ResponseWriter, r *http.Request) {
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusNotFound)
		return
	}

	title, musics_paths, err := album_controller.GetMusicsPathFromAlbum(album_id)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServeZip(w, musics_paths, title+".zip")
}

// downloadAlbumsHandler downloads the albums with the given IDs
// @ Summary: Download albums by their IDs
// @ Description: Download albums by their IDs
// @ Tags: albums
// @ Accept: json
// @ Produces: zip
// @ Param: albums_ids query []int true "Albums Ids"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "albums_ids must be comma separated integers"
// @ Failure 500 {string} string "Unexpected error while downloading albums"
// @ Router /api/albums/download [get]
func downloadAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albums_ids_str := r.URL.Query().Get("albums_ids")
	albums_ids, err := format_utils.ConvertStringToIntArray(albums_ids_str, ",")
	if err != nil {
		http.Error(w, "albums_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	musics_paths, err := album_controller.GetMusicsPathFromAlbums(albums_ids)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}

	file_utils.ServeTreeZip(w, musics_paths, "albums")
}

// getAlbumsByPartialTitleHandler returns the albums with the given partial title
// @ Summary: Get albums by their partial title
// @ Description: Get albums by their partial title
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_partial_title path string true "Album Partial Title"
// @ Success 200 {string} string "OK"
// @ Failure 500 {string} string "Unexpected error while getting albums"
// @ Router /api/albums/partial/{album_partial_title} [get]
func getAlbumsByPartialTitleHandler(w http.ResponseWriter, r *http.Request) {
	album_partial_title := mux.Vars(r)["album_partial_title"]
	albums_json, err := album_controller.GetAlbumsFromPartialTitle(album_partial_title)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.Write(albums_json)
}
