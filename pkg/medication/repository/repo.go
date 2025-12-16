package repository

import (
	"ai_health_assistant/models"
	"gorm.io/gorm"
)
type MedicationRepository interface {
	GetAllMedications(db *gorm.DB) ([]models.Medication, error)
}

