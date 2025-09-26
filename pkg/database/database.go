package database

import (
	"fmt"
	"log"

	"github.com/Sina-Mahmoodmoradi/blog/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)






func Connect(cfg *config.Config)(*gorm.DB, error){

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db,err := gorm.Open(postgres.Open(dns),&gorm.Config{})

	if err!=nil{
		return nil,fmt.Errorf("db connection failed: %w",err)
	}

	log.Println("Connected to DB!")
	
	return db,nil

}