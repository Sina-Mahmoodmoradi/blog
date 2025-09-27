package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type HealthHandler struct {}


func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}


func (h *HealthHandler)RegisterRoutes(r *gin.Engine){
	r.GET("/health",h.HealthCheck)
}

func (h *HealthHandler)HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"ok",
	})
}


