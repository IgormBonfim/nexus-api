package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus/internal/infra/database"
	repository "github.com/igormbonfim/nexus/internal/infra/repositories"
	controller "github.com/igormbonfim/nexus/internal/interfaces/controllers"
	usecase "github.com/igormbonfim/nexus/internal/usecases"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		UserRepository := repository.NewUserRepository(database.DB)
		UserUsecase := usecase.NewUserUsecase(UserRepository)
		UserController := controller.NewUserController(UserUsecase)

		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.POST("/users", UserController.CreateUser)
	}
}
