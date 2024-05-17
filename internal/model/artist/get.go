package artist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// GetArtistsFromFilters returns a list of artists from the database
// Filters can be passed to filter the artists
func GetArtistsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]db_model.Artist, error) {
	var artists []db_model.Artist
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}

// GetArtist returns a artist from the database
// Selects the artist with the given artist_id
func GetArtist(db *gorm.DB, artist_id int) (db_model.Artist, error) {
	var artist db_model.Artist
	// Using `First` to retrieve the first record that matches the artist_id
	result := db.Where("id = ?", artist_id).First(&artist)
	if result.Error != nil {
		return db_model.Artist{}, result.Error
	}
	return artist, nil
}

// GetArtists returns a list of artists from the database
// Selects the artists with the given artist_ids
func GetArtists(db *gorm.DB, artist_ids []int) ([]db_model.Artist, error) {
	var artists []db_model.Artist
	// Using `Find` to retrieve records with the IDs in artist_ids slice
	result := db.Where("id IN ?", artist_ids).Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}

// GetArtistFromPseudo returns a artist from the database.
// Selects the artist with the given pseudo
func GetArtistFromPseudo(db *gorm.DB, pseudo string) (db_model.Artist, error) {
	var artist db_model.Artist
	result := db.Where("pseudo = ?", pseudo).First(&artist)
	if result.Error != nil {
		return db_model.Artist{}, result.Error
	}
	return artist, nil
}

func GetArtistFromPartialPseudo(db *gorm.DB, pseudo string) ([]db_model.Artist, error) {
	var artists []db_model.Artist
	result := db.Where("pseudo LIKE ?", "%"+pseudo+"%").Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}
