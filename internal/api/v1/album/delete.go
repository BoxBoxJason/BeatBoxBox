package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	custom_errors "BeatBoxBox/pkg/errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// deleteAlbumHandler deletes the album with the given ID
// @ Summary: Deletes an album by its ID
// @ Description: Deletes an album by its ID
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album Id"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "album_id must be an integer"
// @ Failure 500 {string} string "Unexpected error while deleting album"
// @ Router /api/albums/{album_id} [delete]
func deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	album_id, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil {
		http.Error(w, "album_id must be an integer", http.StatusBadRequest)
		return
	}

	err = album_controller.DeleteAlbum(album_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteAlbumsHandler deletes the albums with the given IDs
// @ Summary: Deletes albums by their IDs
// @ Description: Deletes albums by their IDs
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: albums_ids query []int true "Albums Ids"
// @ Success 200 {string} string "OK"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 500 {string} string "Unexpected error while deleting albums"
// @ Router /api/albums [delete]
func deleteAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	raw_albums_ids := r.URL.Query().Get("albums_ids")
	albums_ids := []int{}

	if raw_albums_ids != "" {
		albums_ids_str := strings.Split(raw_albums_ids, ",")
		for _, album_id := range albums_ids_str {
			album_id, err := strconv.Atoi(album_id)
			if err != nil {
				http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
				return
			}
			albums_ids = append(albums_ids, album_id)
		}
		err := album_controller.DeleteAlbums(albums_ids)
		if err != nil {
			custom_errors.SendErrorToClient(err, w, "Unexpected error while deleting albums")
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
