package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus/internal/dtos"
	"github.com/igormbonfim/nexus/internal/dtos/requests"
	usecase "github.com/igormbonfim/nexus/internal/usecases"
)

type userController struct {
	validator   dtos.Validator
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase, validator *dtos.Validator) *userController {
	return &userController{
		validator:   *validator,
		userUsecase: *usecase,
	}
}

func (p *userController) CreateUser(ctx *gin.Context) {

	var userDto requests.CreateUserDto
	err := ctx.BindJSON(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.validator.ValidateStruct(&userDto)

	if err != nil {
		formattedErrors := p.validator.FormatValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
		return
	}

	user, err := p.userUsecase.CreateUser(&userDto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
