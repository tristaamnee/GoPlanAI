package main

import (
	"context"
	"goplan-ai/internal/handlers"
	"goplan-ai/internal/repository"
	"goplan-ai/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("goplan_db")
	myMockAI := &services.MockAI{}

	projectRepo := repository.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepo, myMockAI)
	projectHandler := handlers.NewProjectHandler(projectService)

	r := gin.Default()

	r.POST("/projects", projectHandler.CreateProject)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
