package artist_model

type Artist struct {
	Pseudo string `gorm:"type:varchar(32);unique;not null"`
	Id     int    `gorm:"primaryKey;autoIncrement"`
}
