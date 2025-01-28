package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus/internal/domain/entities"
	usecase "github.com/igormbonfim/nexus/internal/usecases"
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

	var user *entities.User
	err := ctx.ShouldBindBodyWithJSON(&user)

	fmt.Println(user)
	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err = p.userUsecase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
