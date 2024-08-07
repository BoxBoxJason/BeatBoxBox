package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	vars := mux.Vars(r)
	album_id, err := strconv.Atoi(vars["album_id"])
	if err != nil || album_id <= 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("album_id is required and must be a positive integer"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	// Delete the album
	err = album_controller.DeleteAlbum(album_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	// Validate album ids
	album_ids, ok := params["id"].([]int)
	if !ok {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Error parsing id, it should be an array of integers"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	// Delete the albums
	err = album_controller.DeleteAlbums(album_ids)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
