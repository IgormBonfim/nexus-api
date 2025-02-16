package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus-api/internal/api/controllers"
	"github.com/igormbonfim/nexus-api/internal/api/middlewares"
	"github.com/igormbonfim/nexus-api/internal/infra/database"
	"github.com/igormbonfim/nexus-api/internal/infra/repositories"
	"github.com/igormbonfim/nexus-api/internal/usecases"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		UserRepository := repositories.NewUserRepository(database.DB)
		UserUsecase := usecases.NewUserUsecase(UserRepository)
		UserController := controllers.NewUserController(UserUsecase)

		api.Use(middlewares.ValidatorMiddleware())

		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.POST("/users", UserController.CreateUser)
		api.POST("/login", UserController.Login)
	}
}
