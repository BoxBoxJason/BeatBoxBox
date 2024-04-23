package album_model

import "gorm.io/gorm"

// PostAlbum is a function that creates a new album in the database.
func PostAlbum(db *gorm.DB, title string, artistId int) error {
	album := Album{
		Title:    title,
		ArtistId: artistId,
	}
	return db.Create(&album).Error
}
