package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil || user_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid user ID, must be a positive integer"))
		return
	}

	err = user_controller.DeleteUser(user_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, nil, nil, nil, []string{"user_id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	if params["user_id"] == nil || len(params["user_id"].([]int)) == 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Missing user_id parameter"))
		return
	}
	err = user_controller.DeleteUsers(params["user_id"].([]int))
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
