package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
)

func postPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	params, err := httputils.ParseMultiPartFormParams(r, []string{"title", "description", "public"}, []string{}, []string{}, []string{"owner_id", "music_id"}, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Validate title
	if params["title"] == nil || params["title"].(string) == "" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Title cannot be empty"))
		return
	}
	// Validate owners ids
	if params["owner_id"] == nil || len(params["owner_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Owner id cannot be empty"))
		return
	}
	title := params["title"].(string)
	description := params["description"].(string)
	public := params["public"].(string) == "true"
	owners_ids := params["owner_id"].([]int)
	musics_ids := params["music_id"].([]int)
	illustration := params["illustration"].(*multipart.FileHeader)
	playlist, err := playlist_controller.PostPlaylist(title, description, public, owners_ids, illustration, musics_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusCreated, playlist)

}
