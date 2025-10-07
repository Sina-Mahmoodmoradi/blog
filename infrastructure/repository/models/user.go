package models

import "time"


type User struct{
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Posts Post `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}