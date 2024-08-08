package music_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// MusicAlreadyExists checks if a music exists in the database
func MusicAlreadyExists(db *gorm.DB, title string, artists_ids []int) bool {
	if len(artists_ids) == 0 {
		return false
	}
	var music db_tables.Music
	err := db.Where("title = ?", title).
		Joins("JOIN artist_musics ON artist_musics.music_id = musics.id").
		Where("artist_musics.artist_id IN ?", artists_ids).
		Group("musics.id").
		Having("COUNT(DISTINCT artist_musics.artist_id) = ?", len(artists_ids)).
		First(&music).Error
	return err == nil

}

// GetMusicsFromFilters returns a list of musics from the database
func GetMusicsFromFilters(db *gorm.DB, titles []string, partial_titles []string, partial_lyrics string, artists []string, albums []string, artist_ids []int, album_ids []int) ([]db_tables.Music, error) {
	query := db.Model(&db_tables.Music{})
	if len(titles) > 0 {
		query = query.Where("title IN ?", titles)
	} else if len(partial_titles) > 0 {
		for _, title := range partial_titles[1:] {
			query = query.Or("title LIKE ?", "%"+title+"%")
		}
	}
	if partial_lyrics != "" {
		query = query.Where("lyrics LIKE ?", "%"+partial_lyrics+"%")
	}
	if len(artist_ids) > 0 {
		query = query.Joins("JOIN artist_musics ON artist_musics.music_id = musics.id").
			Where("artist_musics.artist_id IN ?", artist_ids)
	} else if len(artists) > 0 {
		query = query.Joins("JOIN artist_musics ON artist_musics.music_id = musics.id").
			Joins("JOIN artists ON artists.id = artist_musics.artist_id").
			Where("artists.pseudo IN ?", artists)
	}
	if len(album_ids) > 0 {
		query = query.Joins("JOIN album_musics ON album_musics.music_id = musics.id").
			Where("album_musics.album_id IN ?", album_ids)
	} else if len(albums) > 0 {
		query = query.Joins("JOIN album_musics ON album_musics.music_id = musics.id").
			Joins("JOIN albums ON albums.id = album_musics.album_id").
			Where("albums.title IN ?", albums)
	}
	var musics []db_tables.Music
	err := query.Group("musics.id").Find(&musics).Error
	return musics, err

}

// GetMusic returns a music from the database
func GetMusic(db *gorm.DB, music_id int) (db_tables.Music, error) {
	music := db_model.GetRecordFromId(db, &db_tables.Music{}, music_id)
	if music == nil {
		return db_tables.Music{}, gorm.ErrRecordNotFound
	}
	return *music.(*db_tables.Music), nil
}

// GetMusics returns a list of musics from the database
func GetMusics(db *gorm.DB, music_ids []int) ([]db_tables.Music, error) {
	raw_musics := db_model.GetRecordsFromIds(db, &db_tables.Music{}, music_ids)
	if raw_musics == nil {
		return nil, gorm.ErrRecordNotFound
	}
	musics := make([]db_tables.Music, len(raw_musics))
	for i, music := range raw_musics {
		musics[i] = music.(db_tables.Music)
	}

	return musics, nil
}

// GetMusicsFromPartialTitle returns a list of musics from the database
func GetMusicsFromPartialTitle(db *gorm.DB, title string) []db_tables.Music {
	raw_musics := db_model.GetRecordsByFieldsWithCondition(db, &db_tables.Music{}, map[string]interface{}{}, "title LIKE ?", "%"+title+"%")
	if raw_musics == nil {
		return nil
	}
	musics := make([]db_tables.Music, len(raw_musics))
	for i, music := range raw_musics {
		musics[i] = music.(db_tables.Music)
	}

	return musics
}
