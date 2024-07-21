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

// AddSubscribedPlaylistsToUser adds playlists to the list of subscribed playlists of a user
func AddSubscribedPlaylistsToUser(db *gorm.DB, user *db_tables.User, playlists []*db_tables.Playlist) error {
	return db_model.AddElementsToAssociation(db, user, "SubscribedPlaylists", playlists)
}

// RemoveSubscribedPlaylistsFromUser removes playlists from the list of subscribed playlists of a user
func RemoveSubscribedPlaylistsFromUser(db *gorm.DB, user *db_tables.User, playlists []*db_tables.Playlist) error {
	return db_model.RemoveElementsFromAssociation(db, user, "SubscribedPlaylists", playlists)
}

// AddLikedMusicsToUser adds musics to the list of liked musics of a user
func AddLikedMusicsToUser(db *gorm.DB, user *db_tables.User, musics []*db_tables.Music) error {
	return db_model.AddElementsToAssociation(db, user, "LikedMusics", musics)
}

// RemoveLikedMusicsFromUser removes a music from the list of liked musics of a user
func RemoveLikedMusicsFromUser(db *gorm.DB, user *db_tables.User, musics []*db_tables.Music) error {
	return db_model.RemoveElementsFromAssociation(db, user, "LikedMusics", musics)
}

func RemoveOwnedPlaylistsFromUser(db *gorm.DB, user *db_tables.User, playlists []*db_tables.Playlist) error {
	return db_model.RemoveElementsFromAssociation(db, user, "Playlists", playlists)
}

func AddOwnedPlaylistsToUser(db *gorm.DB, user *db_tables.User, playlists []*db_tables.Playlist) error {
	return db_model.AddElementsToAssociation(db, user, "Playlists", playlists)
}
