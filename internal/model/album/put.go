package album_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

func UpdateAlbum(db *gorm.DB, album_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Updates(update_map).Error
}

func AddMusicsToAlbum(db *gorm.DB, album_id int, musics_ids []int) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Association("Musics").Append(musics_ids)
}

func RemoveMusicsFromAlbum(db *gorm.DB, album_id int, musics_ids []int) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Association("Musics").Delete(musics_ids)
}

func AddArtistToAlbum(db *gorm.DB, album_id int, artist_id int) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Association("Artists").Append(artist_id)
}

func RemoveArtistFromAlbum(db *gorm.DB, album_id int, artist_id int) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Association("Artists").Delete(artist_id)
}
