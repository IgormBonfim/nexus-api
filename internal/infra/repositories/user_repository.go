package repository

import (
	"database/sql"
	"fmt"

	"github.com/igormbonfim/nexus-api/internal/domain/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{
		db: connection,
	}
}

func (repo *UserRepository) InsertUser(user *entities.User) (int, error) {
	var id int
	query, err := repo.db.Prepare("INSERT INTO users" +
		"(username, hashed_password, email, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4, $5) returning id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Username, user.HashedPassword, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}
