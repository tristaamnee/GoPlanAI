package handlers

import (
	"goplan-ai/internal/models"
	"goplan-ai/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func NewProjectHandler(s *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project *models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON Body"})
		return
	}

	err := h.service.CreateProject(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't create project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project created", "data": project})
}
