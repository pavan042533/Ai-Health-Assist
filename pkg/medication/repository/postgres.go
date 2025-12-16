package repository

import(
	"gorm.io/gorm"
	"ai_health_assistant/models"
)

type postgresRepository struct{}

func (r *postgresRepository) GetAllMedications(db *gorm.DB) ([]models.Medication, error) {
	var medications []models.Medication
	result := db.Find(&medications)
	return medications, result.Error
}

func NewMedicationRepository() MedicationRepository {
	return &postgresRepository{}
}