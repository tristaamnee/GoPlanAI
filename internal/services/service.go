package services

import (
	"fmt"
	"goplan-ai/internal/models"
	"goplan-ai/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService struct {
	repo *repository.ProjectRepository
	ai   AIPlanner
}

func NewProjectService(repo *repository.ProjectRepository, ai AIPlanner) *ProjectService {
	return &ProjectService{repo: repo, ai: ai}
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	// AI processor here
	if s.ai == nil {
		return fmt.Errorf("AI engine not initialized")
	}

	project.ID = primitive.NewObjectID()
	project.CreatedAt = time.Now()

	tasks, err := s.ai.GenerateTask(project.Title, project.Description)
	if err == nil {
		project.Tasks = tasks
	}
	return s.repo.Create(*project)
}
