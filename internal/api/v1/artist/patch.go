package artist_handler_v1

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func patchArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve artist id
	artist_id, err := strconv.Atoi(mux.Vars(r)["artist_id"])
	if err != nil || artist_id <= 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("artist_id is required and must be a positive integer"))
		return
	}
	// Parse the query parameters
	params, err := httputils.ParseQueryParams(r, []string{"pseudo", "biography", "birthdate"}, []string{}, []string{}, []string{})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	multipart_params, err := httputils.ParseMultiPartFormParams(r, []string{}, []string{}, []string{}, []string{}, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Validate pseudo
	if params["pseudo"] != nil {
		pseudo, ok := params["pseudo"].(string)
		if ok && pseudo == "" {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("pseudo cannot be empty if provided"))
			return
		}
	}
	// Validate birthdate
	if params["birthdate"] != nil {
		birthdate, ok := params["birthdate"].(string)
		if ok && (birthdate == "" || !format_utils.IsValidDate(birthdate)) {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("birthdate cannot be empty / invalid YYYY-MM-DD format if provided"))
			return
		}
	}
	// Upload illustration if provided
	if multipart_params["illustration"] != nil {
		illustration_file, ok := multipart_params["illustration"].(*multipart.FileHeader)
		if !ok {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("illustration must be a valid file"))
			return
		}
		file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "artist")
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
		params["illustration"] = file_name
	}
	artist, err := artist_controller.UpdateArtist(artist_id, params)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, artist)
}
