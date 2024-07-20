package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
)

func UpdateMusic(music_id int, music_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return err
	}
	return music_model.UpdateMusic(db, &music, music_map)
}

func AddArtistsToMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return err
	}
	artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return err
	} else if artists == nil || len(artists) != len(artists_ids) {
		return custom_errors.NewNotFoundError("some artists were not found")
	}
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return music_model.AddArtistsToMusic(db, &music, artists_ptr)
}

func RemoveArtistsFromMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db.Preload("Artists"), music_id)
	if err != nil {
		return err
	}
	artists_ids_map := make(map[int]bool, len(artists_ids))
	for _, artist_id := range artists_ids {
		artists_ids_map[artist_id] = true
	}
	artists_ptr := make([]*db_tables.Artist, len(artists_ids))
	for i, artist := range music.Artists {
		if _, ok := artists_ids_map[artist.Id]; ok {
			artists_ptr[i] = &artist
		}
	}
	if len(artists_ptr) != len(artists_ids) {
		return custom_errors.NewNotFoundError("some artists were not found")
	}
	return music_model.RemoveArtistsFromMusic(db, &music, artists_ptr)
}
