package usecase

import (
	"github.com/igormbonfim/nexus-api/internal/domain/entities"
	"github.com/igormbonfim/nexus-api/internal/dtos/requests"
	repository "github.com/igormbonfim/nexus-api/internal/infra/repositories"
)

type UserUsecase struct {
	repository *repository.UserRepository
}

func NewUserUsecase(repository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) CreateUser(request *requests.CreateUserDto) (*entities.User, error) {

	user, err := entities.CreateUser(request.Email, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	id, err := u.repository.InsertUser(user)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}
