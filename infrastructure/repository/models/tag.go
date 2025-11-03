package models

import (
	"time"

	"gorm.io/gorm"
)



type Tag struct{
	ID uint `gorm:"primaryKey"`
	Name string
	Posts []Post `gorm:"many2many:post_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}