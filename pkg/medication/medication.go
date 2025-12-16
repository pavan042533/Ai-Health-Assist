package medication

import (
	"ai_health_assistant/pkg/medication/repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type medicationService struct {
	medicationRepository repository.MedicationRepository
}

func NewMedicationService() MedicationService {
	return &medicationService{
		medicationRepository: repository.NewMedicationRepository(),
	}
}

func (s *medicationService) MapMedications(medNames []MedInput, db *gorm.DB) []MappedMedication {
	mappedMeds := make([]MappedMedication, 0)
	for _, med := range medNames {
		product, score, err := FindBestDrugMatchWithAI(db, med.DrugName)
		if err != nil {
			log.Fatal("failed to find best drug match with AI: ", err)
		}
		fmt.Println("product.Ingredients:", product.Ingredients)

		ingIDs := make([]uint, 0)
		for _, pi := range product.Ingredients {
			fmt.Println("pi:",pi)
			fmt.Println("pi.ID:", pi.ID)
			ingIDs = append(ingIDs, pi.ID)
		}

		mappedMeds = append(mappedMeds, MappedMedication{
			DrugName:      product.ProductName,
			IngredientIDs: ingIDs,
			Strength:      med.Strength,
			Frequency:     med.Frequency,
			Duration:      med.Duration,
			MatchScore:    score,
		})
	}
	for _, mappedMed := range mappedMeds {
		fmt.Println(mappedMed.IngredientIDs)
	}
	return mappedMeds
}
