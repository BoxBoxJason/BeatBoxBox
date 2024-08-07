package httputils

import (
	custom_errors "BeatBoxBox/pkg/errors"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, content []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(content)
	if err != nil {
		return custom_errors.NewInternalServerError("Error writing response: " + err.Error())
	}
	return nil
}
