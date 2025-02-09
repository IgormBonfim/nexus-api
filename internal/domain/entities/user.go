package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/igormbonfim/nexus-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"-"`
	PublicKey      string    `json:"public_key"`
	Email          string    `json:"email"`
	Username       string    `json:"nickname"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func CreateUser(email, username, password string) (*User, error) {
	publicKey := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	emailIsValid := utils.IsValidEmail(email)
	if !emailIsValid {
		return nil, errors.New("invalid email format")
	}

	user := &User{
		PublicKey:      publicKey,
		Email:          email,
		Username:       username,
		HashedPassword: string(hashedPassword),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return user, nil
}
