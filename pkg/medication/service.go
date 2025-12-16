package medication

import(
	"gorm.io/gorm"
)

type MedicationService interface {
	MapMedications(medNames []MedInput, db *gorm.DB) []MappedMedication
}