package album_model

import (
	db_tables "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostAlbum is a function that creates a new album in the database.
func CreateAlbum(db *gorm.DB, title string, illustration_path string) (int, error) {
	album := db_tables.Album{
		Title:        title,
		Illustration: illustration_path,
	}
	err := db.Create(&album).Error
	if err != nil {
		return -1, err
	}
	return album.Id, nil
}
