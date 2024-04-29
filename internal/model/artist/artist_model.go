package artist_model

type Artist struct {
	Pseudo       string `gorm:"type:varchar(32);unique;not null"`
	Illustration string `gorm:"type:varchar(36);default:'default.jpg'"`
	Id           int    `gorm:"primaryKey;autoIncrement"`
}
