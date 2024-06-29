package music_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// DeleteMusic deletes an existing music from the database
func DeleteMusic(db *gorm.DB, music_id int) error {
	return db_model.DeleteDBRecord(db, &db_tables.Music{}, music_id)
}

// DeleteMusicFromRecord deletes an existing music from the database
func DeleteMusicFromRecord(db *gorm.DB, music *db_tables.Music) error {
	return db_model.DeleteDBRecordNoFetch(db, music)
}

// DeleteMusics deletes existing musics from the database
func DeleteMusics(db *gorm.DB, music_ids []int) error {
	return db_model.DeleteDBRecords(db, &db_tables.Music{}, music_ids)
}

// DeleteMusicsFromRecords deletes existing musics from the database
func DeleteMusicsFromRecords(db *gorm.DB, musics []*db_tables.Music) error {
	for _, music := range musics {
		err := DeleteMusicFromRecord(db, music)
		if err != nil {
			return err
		}
	}
	return nil
}
