package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"mime/multipart"
	"net/http"
	"strconv"
)

func patchUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL
	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil || user_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid user ID, must be a positive integer"))
		return
	}
	params, err := httputils.ParseMultiPartFormParams(r, []string{"username", "password", "new_password", "email", "bio"}, nil, nil, nil, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if len(params) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("No parameters to update"))
		return
	} else if params["password"] == nil && (params["new_password"] != nil || params["email"] != nil) {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("You must provide your current password to update your email or password"))
		return
	}
	if params["username"] != nil {
		params["pseudo"] = params["username"]
		delete(params, "username")
	}
	// Validate and upload illustration if present
	if params["illustration"] != nil {
		illustration_path, err := file_utils.UploadIllustrationToServer(params["illustration"].(*multipart.FileHeader), "users")
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
		params["illustration"] = illustration_path
	}

	user_json, err := user_controller.UpdateUser(user_id, params)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, user_json)
}

func likeMusicsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get the user ID from the URL
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil || user_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid user ID, must be a positive integer"))
		return
	}
	action := vars["action"]
	if action != "like" && action != "unlike" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid action, must be 'like' or 'unlike'"))
		return
	}
	params, err := httputils.ParseQueryParams(r, nil, nil, nil, []string{"music_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if params["music_id"] == nil || len(params["music_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("You must provide at least 1 music ID"))
		return
	}

	if action == "like" {
		err = user_controller.AddMusicsToLikedMusics(user_id, params["music_id"].([]int))
	} else {
		err = user_controller.RemoveMusicsFromLikedMusics(user_id, params["music_id"].([]int))
	}
}

func subscribePlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get the user ID from the URL
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil || user_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid user ID, must be a positive integer"))
		return
	}
	action := vars["action"]
	if action != "subscribe" && action != "unsubscribe" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid action, must be 'subscribe' or 'unsubscribe'"))
		return
	}
	params, err := httputils.ParseQueryParams(r, nil, nil, nil, []string{"playlist_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	} else if params["playlist_id"] == nil || len(params["playlist_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("You must provide at least 1 playlist ID"))
		return
	}

	if action == "subscribe" {
		err = user_controller.AddPlaylistsToSubscribedPlaylists(user_id, params["playlist_id"].([]int))
	} else {
		err = user_controller.RemovePlaylistsFromSubscribedPlaylists(user_id, params["playlist_id"].([]int))
	}
}
