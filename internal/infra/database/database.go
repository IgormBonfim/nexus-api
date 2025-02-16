package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Banco de dados não está respondendo:", err)
	}

	fmt.Printf("✅ conectado: %s com sucesso!\n", dbname)

	DB = db
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
