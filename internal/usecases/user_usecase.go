package usecase

import (
	"github.com/igormbonfim/nexus/internal/domain/entities"
	"github.com/igormbonfim/nexus/internal/dtos/requests"
	repository "github.com/igormbonfim/nexus/internal/infra/repositories"
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
