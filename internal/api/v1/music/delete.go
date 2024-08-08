package music_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func deleteMusicHandler(w http.ResponseWriter, r *http.Request) {
	music_id, err := strconv.Atoi(mux.Vars(r)["music_id"])
	if err != nil || music_id < 0 {
		custom_errors.SendErrorToClient(w, custom_errors.NewBadRequestError("Invalid music id (must be a positive integer)"))
		return
	}
	err = music_controller.DeleteMusic(music_id)
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteMusicsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, nil, nil, nil, []string{"id"})
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	err = music_controller.DeleteMusics(params["id"].([]int))
	if err != nil {
		custom_errors.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
