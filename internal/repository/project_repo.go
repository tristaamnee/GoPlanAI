package repository

import (
	"context"
	"goplan-ai/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository struct {
	collection *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) *ProjectRepository {
	return &ProjectRepository{
		collection: db.Collection("projects"),
	}
}

func (r *ProjectRepository) Create(project models.Project) error {
	_, err := r.collection.InsertOne(context.Background(), project)
	return err
}
