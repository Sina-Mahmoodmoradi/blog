package middleware

import (
	"net/http"
	"strings"

	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware(tokenManager usecase.TokenManager) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader==""{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"authorization header missing"})
			return 
		}

		parts := strings.Split(authHeader, " ")
		if len(parts)!=2 || parts[0]!="Bearer"{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"invalid authorization header format"})
			return 
		}

		tokenString := parts[1]
		userID,err := tokenManager.ParseToken(tokenString)
		if err!=nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"invalid token"})
			return 
		}

		c.Set("userID",userID)

		c.Next()
	}
}