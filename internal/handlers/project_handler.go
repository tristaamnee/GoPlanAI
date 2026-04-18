package handlers

import (
	"goplan-ai/internal/models"
	"goplan-ai/internal/services"
	"log"
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
	log.Println("--- BẮT ĐẦU CREATE PROJECT ---")
	var project *models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON Body"})
		return
	}
	log.Println("Đang gọi Service để tạo Project...")
	err := h.service.CreateProject(project)
	if err != nil {
		log.Println("SERVICE ERROR:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Println("--- KẾT THÚC THÀNH CÔNG ---")
	c.JSON(http.StatusCreated, gin.H{"message": "Project created", "data": project})
}
