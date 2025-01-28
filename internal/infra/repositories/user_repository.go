package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{
		db: connection,
	}
}
