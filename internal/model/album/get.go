package album_model

import "gorm.io/gorm"

// GetAlbumsFromFilters returns a list of albums from the database
// Filters can be passed to filter the albums
func GetAlbumsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]Album, error) {
	var albums []Album
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&albums)
	if result.Error != nil {
		return nil, result.Error
	}
	return albums, nil
}

// GetAlbum returns a album from the database
// Selects the album with the given album_id
func GetAlbum(db *gorm.DB, album_id int) (Album, error) {
	var album Album
	// Using `First` to retrieve the first record that matches the album_id
	result := db.Where("Id = ?", album_id).First(&album)
	if result.Error != nil {
		return Album{}, result.Error
	}
	return album, nil
}

// GetAlbums returns a list of albums from the database
// Selects the albums with the given album_ids
func GetAlbums(db *gorm.DB, album_ids []int) ([]Album, error) {
	var albums []Album
	// Using `Find` to retrieve records with the IDs in album_ids slice
	result := db.Where("Id IN ?", album_ids).Find(&albums)
	if result.Error != nil {
		return nil, result.Error
	}
	return albums, nil
}
