package models

import (
	"gorm.io/gorm"
)
type Prescription struct {
	gorm.Model
	DoctorName  string
	Date        string    
	ImagePath   string
	Medications []ScannedMedication `gorm:"foreignKey:PrescriptionID"`
}