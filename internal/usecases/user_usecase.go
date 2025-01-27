package usecase

import "github.com/igormbonfim/nexus/internal/domain/entities"

type UserUsecase struct {
	// repository repository.ProductRepository
}

func (u *UserUsecase) CreateUser(user *entities.User) (*entities.User, error) {
	return nil, nil
}
