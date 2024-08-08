package music_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"mime/multipart"
	"net/http"
	"strconv"
)

func patchMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the music_id from the URL
	music_id, err := strconv.Atoi(mux.Vars(r)["music_id"])
	if err != nil || music_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid music_id, must be a positive integer"))
		return
	}
	params, err := httputils.ParseQueryParams(r, []string{"title", "lyrics", "release_date"}, []string{"album_id"}, []string{"genre"}, []string{})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	// Validate title if present
	if params["title"] != nil {
		if len(params["title"].(string)) == 0 {
			custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Title cannot be empty if provided"))
			return
		}
	}
	// Validate date if present
	if params["release_date"] != nil {
		if !format_utils.IsValidDate(params["release_date"].(string)) {
			custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid date format, must be YYYY-MM-DD if provided"))
			return
		}
	}
	multipart_params, err := httputils.ParseMultiPartFormParams(r, []string{}, []string{}, []string{}, []string{}, map[string]string{"illustration": "image", "music": "audio"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	// Validate and upload the illustration if present
	if multipart_params["illustration"] != nil {
		illustration_path, err := file_utils.UploadIllustrationToServer(multipart_params["illustration"].(*multipart.FileHeader), "musics")
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
		params["illustration"] = illustration_path
	}
	// Validate and upload the music if present
	if multipart_params["music"] != nil {
		music_path, err := file_utils.UploadMusicToServer(multipart_params["music"].(*multipart.FileHeader))
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
		params["music"] = music_path
	}

	music_json, err := music_controller.UpdateMusic(music_id, params)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, music_json)
}
