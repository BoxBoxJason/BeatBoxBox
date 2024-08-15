package artist_handler_v1

import (
	artist_controller "BeatBoxBox/internal/controller/artist"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
)

func postArtistHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseMultiPartFormParams(r, []string{"pseudo", "birthdate", "biography"}, []string{}, []string{"genre"}, []string{}, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	// Validate pseudo
	pseudo := ""
	birthdate := ""
	biography := ""
	genres := []string{}
	if params["pseudo"] != nil {
		pseudo, ok := params["pseudo"].(string)
		if ok && pseudo == "" {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("pseudo cannot be empty"))
			return
		}
	}
	// Validate birthdate
	if params["birthdate"] != nil {
		birthdate, ok := params["birthdate"].(string)
		if ok && !format_utils.IsValidDate(birthdate) {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("birthdate must be in the format YYYY-MM-DD"))
			return
		}
	}
	// Validate biography
	biography, _ = params["biography"].(string)
	// Validate genres
	if params["genre"] != nil {
		raw_genres, ok := params["genre"].([]string)
		if !ok {
			httputils.SendErrorToClient(w, httputils.NewBadRequestError("genre must be a list of strings"))
			return
		} else {
			genres = raw_genres
		}
	}
	artist, err := artist_controller.PostArtist(pseudo, genres, biography, birthdate, params["illustration"].(*multipart.FileHeader))
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusCreated, artist)
}
