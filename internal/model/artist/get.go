package artist_model

import "gorm.io/gorm"

// GetArtistsFromFilters returns a list of artists from the database
// Filters can be passed to filter the artists
func GetArtistsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]Artist, error) {
	var artists []Artist
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}

// GetArtist returns a artist from the database
// Selects the artist with the given artist_id
func GetArtist(db *gorm.DB, artist_id int) (Artist, error) {
	var artist Artist
	// Using `First` to retrieve the first record that matches the artist_id
	result := db.Where("Id = ?", artist_id).First(&artist)
	if result.Error != nil {
		return Artist{}, result.Error
	}
	return artist, nil
}

// GetArtists returns a list of artists from the database
// Selects the artists with the given artist_ids
func GetArtists(db *gorm.DB, artist_ids []int) ([]Artist, error) {
	var artists []Artist
	// Using `Find` to retrieve records with the IDs in artist_ids slice
	result := db.Where("Id IN ?", artist_ids).Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}
