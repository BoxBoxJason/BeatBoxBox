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
func GetMusicsFromFilters(db *gorm.DB, filters map[string]interface{}) []db_tables.Music {
	raw_musics := db_model.GetRecordsByFields(db, &db_tables.Music{}, filters)
	if raw_musics == nil {
		return nil
	}
	musics := make([]db_tables.Music, len(raw_musics))
	for i, music := range raw_musics {
		musics[i] = music.(db_tables.Music)
	}

	return musics
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
