package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type HealthHandler struct {
	db *gorm.DB
}


func NewHealthHandler(db *gorm.DB) *HealthHandler{
	return &HealthHandler{
		db:db,
	}
}


func (h *HealthHandler)RegisterRoutes(r *gin.Engine){
	r.GET("/health",h.HealthCheck)
}

func (h *HealthHandler)HealthCheck(c *gin.Context){
	sqlDB,err := h.db.DB()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"db":"not available",
			"error":err.Error(),
		})
	}
	
	if err=sqlDB.Ping();err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"db":"unreachable",
			"error":err.Error(),
		})
	}
	
	
	
	c.JSON(http.StatusOK,gin.H{
		"status":"ok",
		"db":"connected",
	})
}


