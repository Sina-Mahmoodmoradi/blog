package usecase

import "time"



type TokenManager interface {
	CreateToken(userID uint, duration time.Duration)(string,error)
	ParseToken(tokenStr string) (uint ,error)
}