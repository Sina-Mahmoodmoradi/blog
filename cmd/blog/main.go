package main

import (
	"log"

	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http"
	"github.com/Sina-Mahmoodmoradi/blog/pkg/config"
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
	_ = cfg

	r := http.SetupRouter()

	r.Run(":8080")
	
}