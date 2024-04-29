package album_model

import "gorm.io/gorm"

// PostAlbum is a function that creates a new album in the database.
func CreateAlbum(db *gorm.DB, title string, artistId int, illustration_path string) error {
	album := Album{
		Title:        title,
		ArtistId:     artistId,
		Illustration: illustration_path,
	}
	return db.Create(&album).Error
}
