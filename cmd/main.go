package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus-api/internal/api/routes"
	"github.com/igormbonfim/nexus-api/internal/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Iniciando Nexus API")
	env := os.Getenv("APP_ENV")
	log.Println(env)
	if env != "Production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Aviso: Não foi possível carregar .env, usando variáveis de ambiente.")
		}
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Ocorreu um erro ao inicializar o banco de dados: %v", err)
	}

	defer database.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)

	fmt.Println("🚀 Servidor rodando na porta 8000.")
	log.Fatal(server.Run(":8000"))
}
