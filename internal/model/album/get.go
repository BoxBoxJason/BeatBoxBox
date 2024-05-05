package album_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// GetAlbumsFromFilters returns a list of albums from the database
// Filters can be passed to filter the albums
func GetAlbumsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]db_model.Album, error) {
	var albums []db_model.Album
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&albums)
	if result.Error != nil {
		return nil, result.Error
	}
	return albums, nil
}

// GetAlbum returns a album from the database
// Selects the album with the given album_id
func GetAlbum(db *gorm.DB, album_id int) (db_model.Album, error) {
	var album db_model.Album
	// Using `First` to retrieve the first record that matches the album_id
	result := db.Where("id = ?", album_id).First(&album)
	if result.Error != nil {
		return db_model.Album{}, result.Error
	}
	return album, nil
}

// GetAlbums returns a list of albums from the database
// Selects the albums with the given album_ids
func GetAlbums(db *gorm.DB, album_ids []int) ([]db_model.Album, error) {
	var albums []db_model.Album
	// Using `Find` to retrieve records with the IDs in album_ids slice
	result := db.Where("id IN ?", album_ids).Find(&albums)
	if result.Error != nil {
		return nil, result.Error
	}
	return albums, nil
}
