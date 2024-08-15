package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	auth_middleware "BeatBoxBox/internal/middleware/auth"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
)

func postAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	params, err := httputils.ParseMultiPartFormParams(r, []string{"title", "description", "release_date"}, []string{}, []string{"genre"}, []string{"artist_id"}, map[string]string{"illustration": "image/jpeg"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	// Validate artist id
	artists_ids, ok := params["artist_id"].([]int)
	if !ok || len(artists_ids) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Error parsing artist_id, it should have at least one integer"))
	}
	// Validate title
	title, ok := params["title"].(string)
	if !ok || len(title) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Error parsing title, it should be a non empty string"))
	}
	// Validate description
	description, ok := params["description"].(string)
	if !ok || len(description) == 0 {
		description = ""
	}
	// Validate illustration
	illustration, ok := params["illustration"].(*multipart.FileHeader)
	if !ok {
		illustration = nil
	}
	// Validate release date
	release_date, ok := params["release_date"].(string)
	if ok && release_date != "" && !format_utils.IsValidDate(release_date) {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Error parsing release_date, it should be empty OR a valid date in the format YYYY-MM-DD"))
	}

	err = auth_middleware.HasWritePrivileges(r)
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}

	// Create the album
	album, err := album_controller.PostAlbum(title, artists_ids, description, release_date, illustration)
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	httputils.RespondWithJSON(w, http.StatusCreated, album)
}
