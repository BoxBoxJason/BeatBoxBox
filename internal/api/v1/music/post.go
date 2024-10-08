package music_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
)

func postMusicHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseMultiPartFormParams(r, []string{"title", "lyrics", "release_date"}, []string{"album_id"}, []string{"genre"}, []string{"artist_id"}, map[string]string{"illustration": "image", "music": "audio"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
	}
	// Validate title
	if params["title"].(string) == "" {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Title is required and cannot be empty"))
		return
	}
	// Validate release date
	if params["release_date"] != nil && !format_utils.IsValidDate(params["release_date"].(string)) {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid release date format, must be YYYY-MM-DD"))
		return
	}
	// Validate Music file
	if params["music"] == nil {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Music file is required"))
		return
	}
	// Validate artists IDs
	if params["artist_id"] == nil || len(params["artist_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Artist ID is required"))
		return
	}
	title := params["title"].(string)
	lyrics := params["lyrics"].(string)
	release_date := params["release_date"].(string)
	genres := params["genre"].([]string)
	album_id := params["album_id"].(int)
	artists_ids := params["artist_id"].([]int)
	music_file := params["music"].(*multipart.FileHeader)
	illustration_file := params["illustration"].(*multipart.FileHeader)
	music_json, err := music_controller.PostMusic(title, genres, lyrics, release_date, album_id, music_file, illustration_file, artists_ids)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusCreated, music_json)
}
