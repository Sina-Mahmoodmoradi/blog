package models

import (
	"time"

	"gorm.io/gorm"
)







type Comment struct{
	ID uint `gorm:"primaryKey"`
	Content string
	AuthorID uint 
	PostID uint 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}