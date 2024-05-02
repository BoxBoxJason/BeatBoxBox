package album_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostAlbum is a function that creates a new album in the database.
func CreateAlbum(db *gorm.DB, title string, artistId int, illustration_path string) error {
	album := db_model.Album{
		Title:        title,
		Illustration: illustration_path,
	}
	return db.Create(&album).Error
}
