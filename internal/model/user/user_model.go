package user_model

type User struct {
	Pseudo          string `gorm:"type:varchar(32);unique;not null"`
	Email           string `gorm:"type:varchar(256);unique;not null"`
	Hashed_password string `gorm:"type:varchar(64);not null"`
	Salt            string `gorm:"type:varchar(16);not null"`
	Id              int    `gorm:"primaryKey;autoIncrement"`
	Illustration    string `gorm:"type:varchar(36);default:'default.jpg'"`
}
