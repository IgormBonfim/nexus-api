package usecase

import (
	"github.com/igormbonfim/nexus/internal/domain/entities"
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

func (u *UserUsecase) CreateUser(user *entities.User) (*entities.User, error) {
	return nil, nil
}
