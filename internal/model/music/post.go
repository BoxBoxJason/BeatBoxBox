package music_model

import (
	db_tables "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// CreateMusic creates a new music in the database
func CreateMusic(db *gorm.DB, title string, genres []string, lyrics string, album_id int, file_name string, illustration_path string, artists_ptr []*db_tables.Artist) (int, error) {
	if len(artists_ptr) == 0 {
		return -1, nil
	}
	artists := make([]db_tables.Artist, len(artists_ptr))
	for i, artist := range artists_ptr {
		artists[i] = *artist
	}
	new_music := db_tables.Music{
		Title:        title,
		Path:         file_name,
		Genres:       genres,
		Lyrics:       lyrics,
		Illustration: illustration_path,
		Artists:      artists,
	}
	if album_id >= 0 {
		album_id_uint := uint(album_id)
		new_music.AlbumId = &album_id_uint
	}
	err := db.Create(&new_music).Error
	if err != nil {
		return -1, err
	}
	return new_music.Id, nil
}
