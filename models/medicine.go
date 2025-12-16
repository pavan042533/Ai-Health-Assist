package models

type Medication struct {
	ID             uint `gorm:"primaryKey"`
	DrugName       string `json:"drug_name"`
	Amount         string `json:"amount_per_dose"`
	Schedule       string `json:"schedule"`
	Duration       string `json:"duration"`
	Instructions   string `json:"instructions"`
}