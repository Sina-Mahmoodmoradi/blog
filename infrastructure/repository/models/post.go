package models

import "time"



type Post struct{
	ID uint `gorm:"primaryKey"`
	Title string
	Content string
	AuthorID uint 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}