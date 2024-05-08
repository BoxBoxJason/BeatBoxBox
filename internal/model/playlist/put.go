package playlist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdatePlaylist updates an existing playlist in the database
func UpdatePlaylist(db *gorm.DB, playlist_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Playlist{}).Where("id = ?", playlist_id).Updates(update_map).Error
}

func AddMusicsToPlaylist(db *gorm.DB, playlist_id int, music_ids []int) error {
	playlist := db_model.Playlist{}
	err := db.Where("id = ?", playlist_id).First(&playlist).Error
	if err != nil {
		return err
	}
	return db.Model(&playlist).Association("Musics").Append(music_ids)
}

func RemoveMusicsFromPlaylist(db *gorm.DB, playlist_id int, music_ids []int) error {
	playlist := db_model.Playlist{}
	err := db.Where("id = ?", playlist_id).First(&playlist).Error
	if err != nil {
		return err
	}
	return db.Model(&playlist).Association("Musics").Delete(music_ids)
}
