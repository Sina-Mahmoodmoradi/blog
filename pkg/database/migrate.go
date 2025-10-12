package database

import (
	"fmt"
	"log"
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
	"gorm.io/gorm"
)




func AutoMigrate(db *gorm.DB) error{
	err:= db.AutoMigrate(&models.User{},&models.Post{})
	
	if err!=nil{
		return fmt.Errorf("migration failed: %w",err)
	}
	log.Println("AutoMigrate completed")
	return nil
}