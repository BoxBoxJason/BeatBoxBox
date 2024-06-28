package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

func UpdateAlbum(db *gorm.DB, album db_tables.Album, update_map map[string]interface{}) error {
	return db_model.EditRecordFields(db, album, update_map)
}

func AddMusicsToAlbum(db *gorm.DB, album db_tables.Album, musics_ids []int) error {
}

func RemoveMusicsFromAlbum(db *gorm.DB, album_id int, musics_ids []int) error {
}

func AddArtistToAlbum(db *gorm.DB, album_id int, artist_id int) error {
}

func RemoveArtistFromAlbum(db *gorm.DB, album_id int, artist_id int) error {
}
