package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// CreatePlaylist creates a new playlist in the database
func CreatePlaylist(db *gorm.DB, title string, owners_ptr []*db_tables.User, description string, public bool, illustration string, musics_ptr []*db_tables.Music) (db_tables.Playlist, error) {
	owners := make([]db_tables.User, len(owners_ptr))
	for i, owner := range owners_ptr {
		owners[i] = *owner
	}
	musics := make([]db_tables.Music, len(musics_ptr))
	for i, music := range musics_ptr {
		musics[i] = *music
	}
	new_playlist := db_tables.Playlist{
		Title:        title,
		Illustration: illustration,
		Description:  description,
		Public:       public,
		Owners:       owners,
		Subscribers:  owners,
		Musics:       musics,
	}
	err := db.Create(&new_playlist).Error
	if err != nil {
		return db_tables.Playlist{}, err
	}
	return new_playlist, nil
}
