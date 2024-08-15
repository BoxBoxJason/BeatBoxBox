package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	vars := mux.Vars(r)
	album_id, err := strconv.Atoi(vars["album_id"])
	if err != nil || album_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("album_id is required and must be a positive integer"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Delete the album
	err = album_controller.DeleteAlbum(album_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Validate album ids
	album_ids, ok := params["id"].([]int)
	if !ok {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Error parsing id, it should be an array of integers"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Delete the albums
	err = album_controller.DeleteAlbums(album_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
