package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus-api/internal/dtos/requests"
	usecase "github.com/igormbonfim/nexus-api/internal/usecases"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *userController {
	return &userController{
		userUsecase: *usecase,
	}
}

func (p *userController) CreateUser(ctx *gin.Context) {

	validatedData, exists := ctx.Get("validatedData")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}

	userDto, ok := validatedData.(*requests.CreateUserDto)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data type"})
		return
	}

	user, err := p.userUsecase.CreateUser(userDto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
