package dbinit

import (
	db_model "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	"BeatBoxBox/pkg/logger"
)

// Initialize the database connection and create the tables
func init() {
	db, err := db_model.OpenDB()
	if err != nil {
		logger.Critical("Failed to connect database: ", err)
	} else {
		db.AutoMigrate(&artist_model.Artist{}, &user_model.User{}, &album_model.Album{}, &music_model.Music{}, &playlist_model.Playlist{})
		logger.Info("Tables created successfully")
	}
}

// Checks if the database connection is alive
func CheckDB() error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	return nil
}
