package public

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"ai_health_assistant/pkg/medication"
)

func MountRoutes(app *fiber.App) {
	api:= app.Group("/api")

	medicationService := medication.NewMedicationService()
	medicationHandler:= newMedicationHandler(medicationService)

	medicationGroup := api.Group("/medication")
	fmt.Println("Mounting medication routes")
	
	medicationGroup.Post("scan-prescription", medicationHandler.HandleScanPrescription)
	medicationGroup.Post("review-medicine", medicationHandler.HandleReviewMedications)
	// medicationGroup.Post("save", medication.HandleSaveMedications)
	// medicationGroup.Get("list", medication.HandleListMedications)
	// medicationGroup.Post("add", medication.HandleAddMedication)
	// medicationGroup.Put("update/:id", medication.HandleUpdateMedication)
	// medicationGroup.Delete("delete/:id", medication.HandleDeleteMedication)
}