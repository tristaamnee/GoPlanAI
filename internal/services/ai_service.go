package services

import "goplan-ai/internal/models"

type AIPlanner interface {
	GenerateTask(title, description string) ([]models.Task, error)
}

func (g *GroqAI) GenerateTask(title, description string) ([]models.Task, error) {
	return []models.Task{
		{Name: "Thiết kế cấu trúc Database", Priority: "High", Duration: "4h"},
		{Name: "Viết API kết nối AI", Priority: "High", Duration: "8h"},
	}, nil
}
