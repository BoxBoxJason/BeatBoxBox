package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdateMusic updates an existing music in the database
func UpdateMusic(db *gorm.DB, music_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Music{}).Where("id = ?", music_id).Updates(update_map).Error
}

func AddArtistsToMusic(db *gorm.DB, music_id int, artists_ids []int) error {
	return db.Model(&db_model.Music{}).Where("id = ?", music_id).Association("artist_musics").Append(artists_ids)
}

func RemoveArtistsFromMusic(db *gorm.DB, music_id int, artists_ids []int) error {
	return db.Model(&db_model.Music{}).Where("id = ?", music_id).Association("artist_musics").Delete(artists_ids)
}
