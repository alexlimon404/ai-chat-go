package handlers

import (
	"ai-chat-go/database"
	"ai-chat-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AiModelHandler struct{}

func NewAiModelHandler() *AiModelHandler {
	return &AiModelHandler{}
}

// Index GET /go-api/models
func (h *AiModelHandler) Index(c *gin.Context) {
	var aiModels []models.AiModel

	result := database.GetDB().Where("active = ?", true).Find(&aiModels)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch AI models",
		})
		return
	}

	responses := make([]models.AiModelResponse, len(aiModels))
	for i, model := range aiModels {
		responses[i] = model.ToResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
	})
}
