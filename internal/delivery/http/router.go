package http

import (
	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)



func SetupRouter()*gin.Engine{
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()
	healthHandler.RegisterRoutes(r)


	return r
}