package public

import (
	"ai_health_assistant/public/medication"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	api:= app.Group("/api")
	medicationGroup := api.Group("/medication")
	fmt.Println("Mounting medication routes")
	medicationGroup.Post("scan-prescription", medication.HandleScanPrescription)
}
