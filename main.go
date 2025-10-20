package main

import (
	"ai-chat-go/config"
	"ai-chat-go/database"
	"ai-chat-go/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := config.LoadConfig()

	if err := database.Connect(cfg.Database.GetDSN()); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	setupRoutes(router)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(router *gin.Engine) {
	aiModelHandler := handlers.NewAiModelHandler()

	api := router.Group("/go-api")
	{
		api.GET("/models", aiModelHandler.Index)
	}
}
