package db_tables

import (
	"BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	"github.com/lib/pq"
)

// Initialize the database connection and create the tables
func init() {
	db, err := db_model.OpenDB()
	if err != nil {
		logger.Critical("Failed to connect database: ", err)
	} else {
		err = db.AutoMigrate(Artist{}, User{}, Album{}, Music{}, Playlist{}, AuthCookie{})
		if err != nil {
			logger.Info("Tables created successfully")
		} else {
			logger.Critical("Failed to create tables: ", err)
		}
	}
}

type Playlist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Title        string  `gorm:"type:text;unique;not null"`
	Description  string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"many2many:playlist_musics;"`
	CreatorId    int
	Creator      User `gorm:"foreignKey:CreatorId"`
	Protected    bool `gorm:"default:true"`
	CreatedOn    int  `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int  `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Album struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Description  string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:album_artists;"`
	Illustration string   `gorm:"type:text;default:'default.jpg'"`
	Music        []Music  `gorm:"foreignKey:AlbumId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedOn    int      `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int      `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Artist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Pseudo       string  `gorm:"type:varchar(32);unique;not null"`
	Bio          string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"many2many:artist_musics;"`
	CreatedOn    int     `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int     `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Music struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Lyrics       string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:artist_musics;"`
	AlbumId      *uint
	Album        Album          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Genres       pq.StringArray `gorm:"type:text[]"`
	Nblistened   int            `gorm:"default:0"`
	Rating       float32        `gorm:"default:0"`
	Nbrating     int            `gorm:"default:0"`
	Likes        int            `gorm:"default:0"`
	Path         string         `gorm:"type:text;not null"`
	Illustration string         `gorm:"type:text;default:'default.jpg'"`
	UploaderId   *uint
	Uploader     User `gorm:"foreignKey:UploaderId"`
	CreatedOn    int  `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int  `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type User struct {
	Id                  int        `gorm:"primaryKey;autoIncrement"`
	Pseudo              string     `gorm:"type:varchar(32);unique;not null"`
	Email               string     `gorm:"type:varchar(256);unique;not null"`
	Hashed_password     string     `gorm:"type:varchar(64);not null"`
	Illustration        string     `gorm:"type:text;default:'default.jpg'"`
	SubscribedPlaylists []Playlist `gorm:"many2many:user_subscribed_playlists;"`
	Playlists           []Playlist `gorm:"foreignKey:CreatorId"`
	LikedMusics         []Music    `gorm:"many2many:user_liked_musics;"`
	UploadedMusics      []Music    `gorm:"foreignKey:UploaderId"`
	CreatedOn           int        `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn          int        `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type AuthCookie struct {
	Id              int    `gorm:"primaryKey;autoIncrement"`
	HashedAuthToken string `gorm:"type:text;not null"`
	UserId          int
	User            User  `gorm:"foreignKey:UserId"`
	ExpirationDate  int64 `gorm:"not null"`
	CreatedOn       int   `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn      int   `gorm:"autoUpdateTime:milli" json:"modified_on"`
}
