package music_handler_v1

import (
	music_controller "BeatBoxBox/internal/controller/music"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteMusicHandler(w http.ResponseWriter, r *http.Request) {
	music_id, err := strconv.Atoi(mux.Vars(r)["music_id"])
	if err != nil || music_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid music id (must be a positive integer)"))
		return
	}
	err = music_controller.DeleteMusic(music_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteMusicsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, nil, nil, nil, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	err = music_controller.DeleteMusics(params["id"].([]int))
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
