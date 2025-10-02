package tokenmanager

import (
	"errors"
	"time"

	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/golang-jwt/jwt"
)


type JWTManger struct{
	secretKey string
}


func NewJWTTokenManager(secretKey string) usecase.TokenManager{
	return &JWTManger{
		secretKey: secretKey,
	}
}


func (t *JWTManger) CreateToken(userID uint, duration time.Duration)(string,error){
	claims:= jwt.MapClaims{
		"user_id":userID,
		"exp":time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(t.secretKey))
}

func (t *JWTManger)ParseToken(tokenStr string) (uint ,error){
	token, err := jwt.Parse(tokenStr,func(token *jwt.Token)(interface{},error){
		return []byte(t.secretKey),nil
	})

	if err!=nil || !token.Valid{
		return 0,errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return 0,errors.New("invalid claims")
	}

	userID,ok := claims["user_id"].(float64)
	if !ok{
		return 0,errors.New("user_id missing")
	}
	
	return uint(userID),nil

}
