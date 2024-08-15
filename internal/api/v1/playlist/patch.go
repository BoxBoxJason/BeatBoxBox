package playlist_handler_v1

import (
	playlist_controller "BeatBoxBox/internal/controller/playlist"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"mime/multipart"
	"net/http"
	"strconv"
)

func patchPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid playlist id, must be a positive integer"))
		return
	}
	params, err := httputils.ParseQueryParams(r, []string{"title", "description", "release_date", "public"}, nil, nil, nil)
	multipart_form, err := httputils.ParseMultiPartFormParams(r, nil, nil, nil, nil, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	if multipart_form["illustration"] != nil {
		illustration_path, err := file_utils.UploadIllustrationToServer(multipart_form["illustration"].(*multipart.FileHeader), "playlists")
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
		params["illustration"] = illustration_path
	}
	playlist_json, err := playlist_controller.UpdatePlaylist(playlist_id, params)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, playlist_json)
}

func updatePlaylistMusicsHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid playlist id, must be a positive integer"))
		return
	}
	params, err := httputils.ParseQueryParams(r, []string{"action"}, nil, nil, []string{"music_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if len(params["music_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid music_id, must be at least one positive integer"))
		return
	} else if params["action"].(string) != "add" && params["action"].(string) != "remove" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid action, must be 'add' or 'remove'"))
		return
	}

	var playlist_json []byte
	if params["action"].(string) == "add" {
		playlist_json, err = playlist_controller.AddMusicsToPlaylist(playlist_id, params["music_id"].([]int))
	} else {
		playlist_json, err = playlist_controller.RemoveMusicsFromPlaylist(playlist_id, params["music_id"].([]int))
	}

	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, playlist_json)
}

func updatePlaylistOwnersHandler(w http.ResponseWriter, r *http.Request) {
	playlist_id, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid playlist id, must be a positive integer"))
		return
	}
	params, err := httputils.ParseQueryParams(r, []string{"action"}, nil, nil, []string{"user_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if len(params["user_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid user_id, must be at least one positive integer"))
		return
	} else if params["action"].(string) != "add" && params["action"].(string) != "remove" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("invalid action, must be 'add' or 'remove'"))
		return
	}

	var playlist_json []byte
	if params["action"].(string) == "add" {
		playlist_json, err = playlist_controller.AddOwnersToPlaylist(playlist_id, params["user_id"].([]int))
	} else {
		playlist_json, err = playlist_controller.RemoveOwnersFromPlaylist(playlist_id, params["user_id"].([]int))
	}

	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, playlist_json)
}
