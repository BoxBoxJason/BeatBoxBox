package db_tables

import (
	"BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"

	"github.com/lib/pq"
)

func init() {
	CreateTables()
}

// Initialize the database connection and create the tables
func CreateTables() {
	db, err := db_model.OpenDB()
	if err != nil {
		logger.Critical("Failed to connect database: ", err)
	} else {
		if !db.Migrator().HasTable(&Artist{}) {
			logger.Info("Creating tables")
			err = db.AutoMigrate(Artist{}, User{}, Album{}, Music{}, Playlist{}, AuthCookie{}, Role{})
			if err != nil {
				logger.Info("Tables created successfully")
			} else {
				logger.Critical("Failed to create tables: ", err)
			}
		}
	}
}

type Playlist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Title        string  `gorm:"type:text;unique;not null"`
	Description  string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"many2many:playlist_musics;"`
	Owners       []User  `gorm:"many2many:playlists_owners;"`
	Subscribers  []User  `gorm:"many2many:playlists_subscribers;"`
	Protected    bool    `gorm:"default:true"`
	CreatedOn    int     `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int     `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Album struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Description  string   `gorm:"type:text"`
	ReleaseDate  string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:album_artists;"`
	Illustration string   `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music  `gorm:"foreignKey:AlbumId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedOn    int      `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int      `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Artist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Pseudo       string  `gorm:"type:varchar(128);unique;not null"`
	Bio          string  `gorm:"type:text"`
	BirthDate    string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"many2many:artist_musics;"`
	Albums       []Album `gorm:"many2many:album_artists;"`
	CreatedOn    int     `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int     `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type Music struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Lyrics       string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:artist_musics;"`
	ReleaseDate  string   `gorm:"type:text"`
	AlbumId      *uint
	Album        Album          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Genres       pq.StringArray `gorm:"type:text[]"`
	Nblistened   int            `gorm:"default:0"`
	Likes        int            `gorm:"default:0"`
	Path         string         `gorm:"type:text;not null"`
	Illustration string         `gorm:"type:text;default:'default.jpg'"`
	CreatedOn    int            `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn   int            `gorm:"autoUpdateTime:milli" json:"modified_on"`
}

type User struct {
	Id                  int        `gorm:"primaryKey;autoIncrement"`
	Roles               []Role     `gorm:"many2many:user_roles;"`
	Pseudo              string     `gorm:"type:varchar(32);unique;not null"`
	Email               string     `gorm:"type:varchar(256);unique;not null"`
	HashedPassword      string     `gorm:"type:varchar(64);not null"`
	Illustration        string     `gorm:"type:text;default:'default.jpg'"`
	SubscribedPlaylists []Playlist `gorm:"many2many:playlists_subscribers;"`
	Playlists           []Playlist `gorm:"many2many:playlists_owners;"`
	LikedMusics         []Music    `gorm:"many2many:user_liked_musics;"`
	UploadedMusics      int        `gorm:"default:0"`
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

type Role struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(20);unique;not null"`
	Description string `gorm:"type:text"`
	Users       []User `gorm:"many2many:user_roles;"`
}
