package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// GetPlaylistsFromFilters returns a list of playlists from the database
// Filters can be passed to filter the playlists
func GetPlaylistsFromFilters(db *gorm.DB, filters map[string]interface{}) []db_tables.Playlist {
	raw_playlists := db_model.GetRecordsByFields(db, &db_tables.Playlist{}, filters)
	if raw_playlists == nil {
		return nil
	}
	playlists := make([]db_tables.Playlist, len(raw_playlists))
	for i, playlist := range raw_playlists {
		playlists[i] = playlist.(db_tables.Playlist)
	}

	return playlists
}

// GetPlaylist returns a playlist from the database
// Selects the playlist with the given playlist_id
func GetPlaylist(db *gorm.DB, playlist_id int) (db_tables.Playlist, error) {
	playlist := db_model.GetRecordFromId(db, &db_tables.Playlist{}, playlist_id)
	if playlist == nil {
		return db_tables.Playlist{}, gorm.ErrRecordNotFound
	}
	return *playlist.(*db_tables.Playlist), nil
}

// GetPlaylists returns a list of playlists from the database
// Selects the playlists with the given playlist_ids
func GetPlaylists(db *gorm.DB, playlist_ids []int) ([]db_tables.Playlist, error) {
	records := db_model.GetRecordsFromIds(db, &db_tables.Playlist{}, playlist_ids)
	if records == nil {
		return nil, gorm.ErrRecordNotFound
	}
	playlists := make([]db_tables.Playlist, len(records))
	for i, record := range records {
		playlists[i] = record.(db_tables.Playlist)
	}

	return playlists, nil
}

func GetPlaylistsFromPartialTitle(db *gorm.DB, filters map[string]interface{}, title string) []db_tables.Playlist {
	records := db_model.GetRecordsByFieldsWithCondition(db, &db_tables.Playlist{}, filters, "title LIKE ?", "%"+title+"%")
	if records == nil {
		return nil
	}
	playlists := make([]db_tables.Playlist, len(records))
	for i, record := range records {
		playlists[i] = record.(db_tables.Playlist)
	}

	return playlists
}
