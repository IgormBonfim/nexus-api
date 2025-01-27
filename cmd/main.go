package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igormbonfim/nexus/internal/infra/database"
	"github.com/igormbonfim/nexus/internal/interfaces/routes"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Ocorreu um erro ao inicializar o banco de dados: %v", err)
	}

	defer database.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8000")
}
