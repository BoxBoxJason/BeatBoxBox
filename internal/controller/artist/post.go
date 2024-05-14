package artist_controller

import (
	db_model "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	custom_errors "BeatBoxBox/pkg/errors"
)

func PostArtist(pseudo string, illustration_file_name string) (int, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return -1, err
	}
	defer db_model.CloseDB(db)

	// Check if the pseudo is empty or already taken
	if pseudo == "" {
		return -1, custom_errors.NewBadRequestError("artist pseudo is empty")
	}
	if IsPseudoTaken(pseudo) {
		return -1, custom_errors.NewBadRequestError("artist already exists")
	}

	return artist_model.CreateArtist(db, pseudo, illustration_file_name)
}
