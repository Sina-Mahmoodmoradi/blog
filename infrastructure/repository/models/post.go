package models

import (
	"time"

	"gorm.io/gorm"
)



type Post struct{
	ID uint `gorm:"primaryKey"`
	Title string
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string
	AuthorID uint 
	Tags []Tag `gorm:"many2many:post_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}