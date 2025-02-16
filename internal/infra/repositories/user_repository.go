package repositories

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
		"(public_key, username, hashed_password, email, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4, $5, $6) returning id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.PublicKey, user.Username, user.HashedPassword, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	queryString := "SELECT * FROM users WHERE users.email = $1"

	query, err := repo.db.Prepare(queryString)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user entities.User

	err = query.QueryRow(email).Scan(
		&user.ID,
		&user.PublicKey,
		&user.Username,
		&user.HashedPassword,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &user, nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*entities.User, error) {
	queryString := "SELECT * FROM users WHERE users.username = $1"

	query, err := repo.db.Prepare(queryString)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user entities.User

	err = query.QueryRow(username).Scan(
		&user.ID,
		&user.PublicKey,
		&user.Username,
		&user.HashedPassword,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &user, nil
}
