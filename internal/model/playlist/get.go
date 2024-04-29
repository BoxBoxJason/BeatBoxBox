package playlist_model

import "gorm.io/gorm"

// GetPlaylistsFromFilters returns a list of playlists from the database
// Filters can be passed to filter the playlists
func GetPlaylistsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]Playlist, error) {
	var playlists []Playlist
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&playlists)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlists, nil
}

// GetPlaylist returns a playlist from the database
// Selects the playlist with the given playlist_id
func GetPlaylist(db *gorm.DB, playlist_id int) (Playlist, error) {
	var playlist Playlist
	// Using `First` to retrieve the first record that matches the playlist_id
	result := db.Where("Id = ?", playlist_id).First(&playlist)
	if result.Error != nil {
		return Playlist{}, result.Error
	}
	return playlist, nil
}

// GetPlaylists returns a list of playlists from the database
// Selects the playlists with the given playlist_ids
func GetPlaylists(db *gorm.DB, playlist_ids []int) ([]Playlist, error) {
	var playlists []Playlist
	// Using `Find` to retrieve records with the IDs in playlist_ids slice
	result := db.Where("Id IN ?", playlist_ids).Find(&playlists)
	if result.Error != nil {
		return nil, result.Error
	}
	return playlists, nil
}