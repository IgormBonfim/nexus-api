package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/igormbonfim/nexus-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrEmailTooLong       = errors.New("email cannot exceed 255 characters")
	ErrUsernameTooLong    = errors.New("username cannot exceed 50 characters")
	ErrUsernameTooShort   = errors.New("username must have at least 3 characters")
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

	if err := validateUsername(username); err != nil {
		return nil, err
	}

	if err := validateEmail(email); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := &User{
		PublicKey:      publicKey,
		Email:          email,
		Username:       username,
		HashedPassword: string(hashedPassword),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	return user, nil
}

func validateUsername(username string) error {
	len := len(username)

	if len > 50 {
		return ErrUsernameTooLong
	}

	if len < 3 {
		return ErrUsernameTooShort
	}

	return nil
}

func validateEmail(email string) error {
	len := len(email)

	if len > 50 {
		return ErrEmailTooLong
	}

	emailIsValid := utils.IsValidEmail(email)
	if !emailIsValid {
		return ErrInvalidEmailFormat
	}

	return nil
}
