package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


type Config struct {
	AppEnv       string
	Port         string
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JWTSecret    string
}


func LoadFromEnv()(*Config,error){

	get:= func (key,def string) string {
		if v,ok:=os.LookupEnv(key);ok && v!=""{
			return v
		}
		return def
	}

	dbPort,err := strconv.Atoi(get("DB_PORT","5432"))
	if err!=nil{
		return nil,fmt.Errorf("invalid DB_PORT: %w",err)
	}

	readTimeOut, err := strconv.Atoi(get("READ_TIMEOUT_SEC","15"))
	if err!=nil{
		return nil,fmt.Errorf("invalid READ_TIMEOUT_SEC: %w",err)
	}


	writeTimeOut, err := strconv.Atoi(get("WRITE_TIMEOUT_SEC","15"))
	if err!=nil{
		return nil,fmt.Errorf("invalid WRITE_TIMEOUT_SEC: %w",err)
	}

	cfg := &Config{
		AppEnv:       get("APP_ENV","development"),
		Port:         get("PORT","8080"),
		DBHost:       get("DB_HOST","localhost"),
		DBPort:       dbPort,
		DBUser:       get("DB_USER","postgres"),
		DBPassword:   get("DB_PASSWORD",""),
		DBName:       get("DB_NAME","app"),
		ReadTimeout:  time.Duration(readTimeOut) * time.Second,
		WriteTimeout: time.Duration(writeTimeOut) * time.Second,
		JWTSecret:    get("JWT_SECRET",""),
	}

	if cfg.DBPassword==""{
		return nil,fmt.Errorf("DB_PASSWORD is required")
	}


	if cfg.JWTSecret==""{
		return nil,fmt.Errorf("JWT_SECRET is required")
	}

	return cfg,nil

}