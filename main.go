package main

import (
	"log"

	"github.com/Andresito126/api3-notifications/src/infrastructure/dependencies"
	"github.com/Andresito126/api3-notifications/src/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// dependencias
	dependencies.InitDependencies()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8082")
}