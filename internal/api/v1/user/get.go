package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	"BeatBoxBox/pkg/utils/httputils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user id from the URL
	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil || user_id < 0 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("Invalid user id, must be a positive integer"))
		return
	}
	user_json, err := user_controller.GetUserJSON(user_id)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, user_json)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	params, err := httputils.ParseQueryParams(r, nil, nil, []string{"pseudo", "partial_pseudo"}, []string{"id"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	if len(params) != 1 {
		httputils.SendErrorToClient(w, httputils.NewBadRequestError("You must provide exactly one parameter (id, pseudo or partial_pseudo)"))
		return
	}
	var users_json []byte
	if params["id"] != nil {
		users_json, err = user_controller.GetUsersJSON(params["id"].([]int))
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	} else if params["pseudo"] != nil {
		users_json, err = user_controller.GetUsersFromPseudos(params["pseudo"].([]string))
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	} else if params["partial_pseudo"] != nil {
		users_json, err = user_controller.GetUsersFromPartialPseudos(params["partial_pseudo"].([]string))
		if err != nil {
			httputils.SendErrorToClient(w, err)
			return
		}
	}

	httputils.RespondWithJSON(w, http.StatusOK, users_json)
}
