package main

import (
	"goplan-ai/internal/handlers"
	"goplan-ai/internal/repository"
	"goplan-ai/internal/services"
	"goplan-ai/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := database.ConnectDB(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(os.Getenv("DB_NAME"))

	groqAPIKey := os.Getenv("GROQ_API_KEY")
	if groqAPIKey == "" {
		log.Fatal("GROQ_API_KEY environment variable not set")
	}

	aiEngine := &services.GroqAI{
		APIKey: os.Getenv(groqAPIKey),
	}

	projectRepo := repository.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepo, aiEngine)
	projectHandler := handlers.NewProjectHandler(projectService)

	r := gin.Default()

	r.POST("/projects", projectHandler.CreateProject)

	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
