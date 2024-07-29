package album_model

import (
	db_tables "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostAlbum is a function that creates a new album in the database.
func CreateAlbum(db *gorm.DB, title string, description string, illustration_path string, artists_ptr []*db_tables.Artist) (int, error) {
	artists := make([]db_tables.Artist, len(artists_ptr))
	for i, artist := range artists_ptr {
		artists[i] = *artist
	}
	album := db_tables.Album{
		Title:        title,
		Illustration: illustration_path,
		Description:  description,
		Artists:      artists,
	}
	err := db.Create(&album).Error
	if err != nil {
		return -1, err
	}
	return album.Id, nil
}
