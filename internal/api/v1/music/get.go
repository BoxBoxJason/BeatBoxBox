package music_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL
	music_id, err := strconv.Atoi(mux.Vars(r)["music_id"])
	if err != nil || music_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid music ID, must be a positive integer"))
		return
	}
	music, err := music_controller.GetMusicJSON(music_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, music)
}

func getMusicsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, []string{"partial_lyrics"}, []string{}, []string{"title", "partial_title", "artist", "album"}, []string{"music_id", "artist_id", "album_id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	var musics []byte
	if params["music_id"] != nil && len(params) > 1 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid query parameters: music_id cannot be used with other parameters"))
		return
	} else if params["music_id"] != nil {
		musics, err = music_controller.GetMusicsJSON(params["music_id"].([]int))
		if err != nil {
			custom_errors.SendErrorToClient(w, err)
			return
		}
	} else {
		titles := params["title"].([]string)
		partial_titles := params["partial_title"].([]string)
		partial_lyrics := params["partial_lyrics"].(string)
		artists := params["artist"].([]string)
		albums := params["album"].([]string)
		artist_ids := params["artist_id"].([]int)
		album_ids := params["album_id"].([]int)
		musics, err = music_controller.GetMusicsJSONFromFilters(titles, partial_titles, partial_lyrics, artists, albums, artist_ids, album_ids)
	}
	httputils.RespondWithJSON(w, http.StatusOK, musics)
}

func downloadMusicHandler(w http.ResponseWriter, r *http.Request) {
	music_id, err := strconv.Atoi(mux.Vars(r)["music_id"])
	if err != nil || music_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid music ID, must be a positive integer"))
		return
	}
	err = music_controller.ServeMusicFile(w, music_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
	}
}

func downloadMusicsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, []string{}, []string{}, []string{}, []string{"id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	if params["id"] == nil {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("id is required"))
		return
	}
	err = music_controller.ServeMusicsFiles(w, params["id"].([]int))
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
	}
}
