package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	music_controller "BeatBoxBox/internal/controller/music"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func patchAlbumHandler(w http.ResponseWriter, r *http.Request) {
	album_id, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil || album_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("album_id is required and must be a positive integer"))
		return
	}
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{"title", "description", "release_date"}, []string{}, []string{"genre"}, []string{})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	// Validate title
	title, ok := params["title"].(string)
	if ok && title == "" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("title cannot be empty if provided"))
		return
	}
	// Validate release date
	release_date, ok := params["release_date"].(string)
	if ok && (release_date == "" || !format_utils.IsValidDate(release_date)) {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("release_date cannot be empty / invalid YYYY-MM-DD format if provided"))
		return
	}
	body_params, err := httputils.ParseMultiPartFormParams(r, []string{}, []string{}, []string{}, []string{}, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	if body_params["illustration"] != nil {
		illustration_file, ok := body_params["illustration"].(*multipart.FileHeader)
		if !ok {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("illustration must be a valid file"))
			return
		}
		file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "album")
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
		params["illustration"] = file_name
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	album, err := album_controller.UpdateAlbum(album_id, params)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, album)
}

func patchAlbumArtistsHandler(w http.ResponseWriter, r *http.Request) {
	album_id, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil || album_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("album_id is required and must be a positive integer"))
		return
	}
	action := mux.Vars(r)["action"]
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"artist_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	artists_ids, ok := params["artist_id"].([]int)
	if !ok || len(artists_ids) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("artist_id is required and must contain at least one positive integer"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	album, err := album_controller.UpdateAlbumArtists(album_id, artists_ids, action)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, album)
}

func patchAlbumMusicsHandler(w http.ResponseWriter, r *http.Request) {
	album_id, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil || album_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("album_id is required and must be a positive integer"))
		return
	}
	action := mux.Vars(r)["action"]
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"music_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	musics_ids, ok := params["music_id"].([]int)
	if !ok || len(musics_ids) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("music_id is required and must contain at least one positive integer"))
		return
	}
	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	var album []byte
	if action == "add" {
		album, err = music_controller.AddMusicsToAlbum(musics_ids, album_id)
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	} else if action == "remove" {
		album, err = music_controller.RemoveAlbumFromMusics(musics_ids, album_id)
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	} else {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("action must be 'add' or 'remove'"))
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, album)
}
