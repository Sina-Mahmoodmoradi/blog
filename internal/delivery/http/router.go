package http

import "github.com/gin-gonic/gin"



func SetupRouter()*gin.Engine{
	r := gin.Default()

	r.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"pong",
		})
	})
	


	return r
}