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

func (u *UserUsecase) CreateUser(user *requests.CreateUserDto) (*entities.User, error) {
	return nil, nil
}
