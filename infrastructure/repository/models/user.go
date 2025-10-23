package models

import (
	"time"

	"gorm.io/gorm"
)


type User struct{
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Posts []Post `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}