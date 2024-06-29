package user_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// UpdateUser updates an existing user in the database
func UpdateUser(db *gorm.DB, user *db_tables.User, fields map[string]interface{}) error {
	return db_model.EditRecordFields(db, user, fields)
}

// AddSubscribedPlaylistToUser adds a playlist to the list of subscribed playlists of a user
func AddSubscribedPlaylistToUser(db *gorm.DB, user *db_tables.User, playlist *db_tables.Playlist) error {
	return db.Model(user).Association("SubscribedPlaylists").Append(playlist)
}

// RemoveSubscribedPlaylistFromUser removes a playlist from the list of subscribed playlists of a user
func RemoveSubscribedPlaylistFromUser(db *gorm.DB, user *db_tables.User, playlist *db_tables.Playlist) error {
	return db.Model(user).Association("SubscribedPlaylists").Delete(playlist)
}

// AddLikedMusicToUser adds a music to the list of liked musics of a user
func AddLikedMusicToUser(db *gorm.DB, user *db_tables.User, music *db_tables.Music) error {
	return db.Model(user).Association("LikedMusics").Append(music)
}

// RemoveLikedMusicFromUser removes a music from the list of liked musics of a user
func RemoveLikedMusicFromUser(db *gorm.DB, user *db_tables.User, music *db_tables.Music) error {
	return db.Model(user).Association("LikedMusics").Delete(music)
}
