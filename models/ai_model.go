package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AiModel struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UUID        string    `gorm:"type:uuid;uniqueIndex;not null" json:"uuid"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Group       string    `gorm:"type:varchar(255);not null" json:"group"`
	Task        string    `gorm:"type:varchar(255);not null" json:"task"`
	Driver      string    `gorm:"type:varchar(255);not null" json:"driver"`
	ModelName   string    `gorm:"type:varchar(255);not null;column:model_name" json:"model_name"`
	Host        string    `gorm:"type:varchar(255)" json:"host,omitempty"`
	Token       string    `gorm:"type:text" json:"token,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Active      bool      `gorm:"default:true" json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (AiModel) TableName() string {
	return "ai_models"
}

func (m *AiModel) BeforeCreate(tx *gorm.DB) error {
	if m.UUID == "" {
		m.UUID = uuid.New().String()
	}
	return nil
}

const (
	GroupNaturalLanguageProcessing = "natural_language_processing"
)

const (
	TaskTextGeneration = "text_generation"
)

const (
	DriverLocalAI = "localAi"
	DriverTest    = "test"
)

type AiModelResponse struct {
	ID        uint   `json:"id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Task      string `json:"task"`
	Driver    string `json:"driver"`
	ModelName string `json:"model_name"`
	Active    bool   `json:"active"`
}

func (m *AiModel) ToResponse() AiModelResponse {
	return AiModelResponse{
		ID:        m.ID,
		UUID:      m.UUID,
		Name:      m.Name,
		Group:     m.Group,
		Task:      m.Task,
		Driver:    m.Driver,
		ModelName: m.ModelName,
		Active:    m.Active,
	}
}
