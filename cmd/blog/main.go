package main

import (
	"log"

	delivery "github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http"
	"github.com/Sina-Mahmoodmoradi/blog/pkg/config"
	"github.com/Sina-Mahmoodmoradi/blog/pkg/database"
	"github.com/joho/godotenv"
)


func main(){

	if err:=godotenv.Load();err!=nil{
		log.Println("no .env file found (continuing with system env)")
	}

	cfg,err:= config.LoadFromEnv()
	if err!=nil{
		log.Fatalf("failed to load config: %v",err)
	}
	
	db,err := database.Connect(cfg)
	if err!=nil{
		log.Fatalf("database connection error: %v",err)
	}

	if err:=database.AutoMigrate(db);err!=nil{
		log.Fatalf("migration failed: %v",err)
	}

	r := delivery.SetupRouter(db)

	r.Run(":8080")
	
}