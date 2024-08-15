package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	"BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

// DeleteMusic deletes a music by its id
func DeleteMusic(music_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return err
	}
	return music_model.DeleteMusicFromRecord(db, &music)
}

// DeleteMusics deletes musics by their ids
func DeleteMusics(music_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db, music_ids)
	if err != nil {
		return err
	} else if musics == nil || len(musics) != len(music_ids) {
		return httputils.NewNotFoundError("some musics were not found")
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	return music_model.DeleteMusicsFromRecords(db, musics_ptr)
}
