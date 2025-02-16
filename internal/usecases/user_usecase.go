package usecases

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/igormbonfim/nexus-api/internal/domain/entities"
	"github.com/igormbonfim/nexus-api/internal/dtos/requests"
	"github.com/igormbonfim/nexus-api/internal/infra/repositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrRegisteredEmail    = errors.New("this email address is already registered. please use a different one")
	ErrRegisteredUsername = errors.New("this username is already registered. please use a different one")
	ErrInvalidCredentials = errors.New("please check your username and password and try again")
)

type UserUsecase struct {
	repository *repositories.UserRepository
}

func NewUserUsecase(repository *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) CreateUser(request *requests.CreateUserDto) (*entities.User, error) {

	user, err := entities.CreateUser(request.Email, request.Username, request.Password)
	if err != nil {
		return nil, fmt.Errorf("user creation failed: %w", err)
	}

	userExists, err := u.repository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if userExists != nil {
		return nil, ErrRegisteredEmail
	}

	userExists, err = u.repository.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if userExists != nil {
		return nil, ErrRegisteredUsername
	}

	id, err := u.repository.InsertUser(user)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (u *UserUsecase) LoginUser(request *requests.LoginDto) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	user, err := u.repository.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrInvalidCredentials
	}

	fmt.Println(user.HashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password))
	if err != nil {
		fmt.Println(err)
		return "", ErrInvalidCredentials
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := jwt.RegisteredClaims{
		ID:        uuid.NewString(),
		Subject:   user.PublicKey,
		Issuer:    "Nexus-api",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
