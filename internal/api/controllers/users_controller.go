package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus-api/internal/domain/entities"
	"github.com/igormbonfim/nexus-api/internal/dtos/requests"
	"github.com/igormbonfim/nexus-api/internal/dtos/responses"
	"github.com/igormbonfim/nexus-api/internal/usecases"
)

type userController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(usecase *usecases.UserUsecase) *userController {
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
		if errors.Is(err, usecases.ErrRegisteredEmail) || errors.Is(err, entities.ErrInvalidEmailFormat) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (p *userController) Login(ctx *gin.Context) {
	validatedData, exists := ctx.Get("validatedData")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}

	loginDto, ok := validatedData.(*requests.LoginDto)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data type"})
		return
	}

	token, err := p.userUsecase.LoginUser(loginDto)
	if err != nil {
		if err == usecases.ErrInvalidCredentials {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.NewLoginResponse(token))
}
