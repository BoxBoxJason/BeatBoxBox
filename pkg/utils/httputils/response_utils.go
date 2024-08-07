package httputils

import (
	custom_errors "BeatBoxBox/pkg/errors"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(content)
	if err != nil {
		custom_errors.SendErrorToClient(w, custom_errors.NewInternalServerError("Error writing response: "+err.Error()))
	}
}
