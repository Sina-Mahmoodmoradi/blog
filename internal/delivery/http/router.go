package http

import (
	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func SetupRouter(db *gorm.DB)*gin.Engine{
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()
	healthHandler.RegisterRoutes(r)


	return r
}