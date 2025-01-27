package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
