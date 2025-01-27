package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func Connect() error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("Erro ao conectar com o banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Printf("Conectado: %s", dbname)

	DB = db
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
