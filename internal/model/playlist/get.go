package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// GetPlaylistsFromFilters returns a list of playlists from the database
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
func GetPlaylist(db *gorm.DB, playlist_id int) (db_tables.Playlist, error) {
	playlist := db_model.GetRecordFromId(db, &db_tables.Playlist{}, playlist_id)
	if playlist == nil {
		return db_tables.Playlist{}, gorm.ErrRecordNotFound
	}
	return *playlist.(*db_tables.Playlist), nil
}

// GetPlaylists returns a list of playlists from the database
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

func PlaylistAlreadyExists(db *gorm.DB, playlist_name string, owners_ids []int) bool {
	if len(owners_ids) == 0 {
		return false
	}
	var album db_tables.Playlist
	err := db.Preload("Owners").Where("title = ?", playlist_name).
		Joins("JOIN playlists_owners ON playlists_owners.playlist_id = playlists.id").
		Where("playlists_owners.user_id IN ?", owners_ids).
		Group("playlists.id").
		Having("COUNT(DISTINCT playlists_owners.user_id) = ?", len(owners_ids)).
		First(&album).Error
	return err == nil
}

func GetPlaylistsByFilters(db *gorm.DB, titles []string, musics []string, owners []string, music_ids []int, owner_ids []int) ([]db_tables.Playlist, error) {
	query := db.Model(&db_tables.Playlist{})
	if len(titles) > 0 {
		query = query.Where("title IN ?", titles)
	}
	if len(music_ids) > 0 {
		query = query.Joins("JOIN playlist_musics ON playlist_musics.playlist_id = playlists.id").
			Where("playlist_musics.music_id IN ?", music_ids)
	} else if len(musics) > 0 {
		query = query.Joins("JOIN playlist_musics ON playlist_musics.playlist_id = playlists.id").
			Joins("JOIN musics ON musics.id = playlist_musics.music_id").
			Where("musics.title IN ?", musics)
	}
	if len(owner_ids) > 0 {
		query = query.Joins("JOIN playlists_owners ON playlists_owners.playlist_id = playlists.id").
			Where("playlists_owners.user_id IN ?", owner_ids)
	} else if len(owners) > 0 {
		query = query.Joins("JOIN playlists_owners ON playlists_owners.playlist_id = playlists.id").
			Joins("JOIN users ON users.id = playlists_owners.user_id").
			Where("users.pseudo IN ?", owners)
	}
	var playlists []db_tables.Playlist
	err := query.Group("playlists.id").Find(&playlists).Error
	return playlists, err
}
