package user_handler_v1

import (
	user_controller "BeatBoxBox/internal/controller/user"
	"BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
	"net/http"
)

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	params, err := httputils.ParseMultiPartFormParams(r, []string{"username", "email", "password", "bio"}, nil, nil, nil, map[string]string{"illustration": "image"})
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	user_json, err := user_controller.PostUser(params["username"].(string), params["email"].(string), params["password"].(string), params["bio"].(string), params["illustration"].(*multipart.FileHeader))
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.RespondWithJSON(w, http.StatusCreated, user_json)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	params, err := httputils.ParseFormBodyParams(r, []string{"email", "password", "username", "username_or_email"}, []string{"id"}, nil, nil)
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	var user_id int
	var raw_session_token string
	if len(params) < 2 || params["password"] == nil {
		httputils.SendErrorToClient(w, httputils.NewUnauthorizedError("Authentication failed"))
		return
	} else if params["email"] != nil {
		user_id, raw_session_token, err = user_controller.AttemptLoginFromEmail(params["email"].(string), params["password"].(string))
	} else if params["username"] != nil {
		user_id, raw_session_token, err = user_controller.AttemptLoginFromUsername(params["username"].(string), params["password"].(string))
	} else if params["id"] != nil {
		user_id, raw_session_token, err = user_controller.AttemptLoginFromId(params["id"].(int), params["password"].(string))
	} else {
		user_id, raw_session_token, err = user_controller.AttemptLogin(params["username_or_email"].(string), params["password"].(string))
	}
	if err != nil {
		httputils.SendErrorToClient(w, err)
		return
	}
	httputils.SetAuthCookie(w, user_id, raw_session_token)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	httputils.SetAuthCookie(w, -1, "")
	http.Redirect(w, r, "/", http.StatusFound)
}
