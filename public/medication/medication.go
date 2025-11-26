package medication

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func HandleScanPrescription(c *fiber.Ctx) error {
	// 1. Parse image
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded. Use -F 'image=@file.jpg'"})
	}

	file, _ := fileHeader.Open()
	defer file.Close()
	imgData, _ := io.ReadAll(file)

	mimeType := http.DetectContentType(imgData)
	var format string
	switch mimeType {
	case "image/jpeg":
		format = "jpeg"
	case "image/png":
		format = "png"
	case "image/webp":
		format = "webp"
	case "image/heic":
		format = "heic"
	default:
		// Fallback for unknown types, often handled well by Gemini as JPEG
		format = "jpeg"
		log.Printf("Warning: Unknown mime type %s, defaulting to jpeg", mimeType)
	}

	// 2. Check API Key
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "GEMINI_API_KEY is missing in server environment"})
	}

	// 3. Setup Client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Client create failed: " + err.Error()})
	}
	defer client.Close()

	// 4. Call Gemini
	model := client.GenerativeModel("gemini-2.0-flash")
	model.ResponseMIMEType = "application/json"

	prompt := `
			You are a medical assistant. Extract medication details from the following OCR text of a prescription.
			Output strictly valid JSON. No markdown, no explanations.

			OCR TEXT:

			JSON SCHEMA:
			{
			"medications": [
				{
				"drug_name": "string",
				"amount_per_dose": "string (e.g, 1 tab, \"3 ml\", etc.)",
				"schedule": "string (should be : 'Once daily', \"Twice daily\", \"Three times daily\ As needed, etc.)",)",
				"duration": "string",
				"instructions": "string"
				}
			]
			}
			`
	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt),
		genai.ImageData(format, imgData),
	)

	// *** DEBUGGING: Return the RAW error if it fails ***
	if err != nil {
		log.Println("Gemini Error:", err)
		return c.Status(500).JSON(fiber.Map{
			"error":   "Gemini API Error",
			"details": err.Error(),
		})
	}

	// --- STEP 6: EXTRACT & CLEAN RESPONSE ---
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "AI returned no content."})
	}

	var jsonResponseString string
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			jsonResponseString += string(txt)
		}
	}

	// Remove Markdown fences if present (```json ... ```)
	jsonResponseString = strings.TrimPrefix(jsonResponseString, "```json")
	jsonResponseString = strings.TrimPrefix(jsonResponseString, "```")
	jsonResponseString = strings.TrimSuffix(jsonResponseString, "```")

	// Trim whitespace to ensure clean JSON
	jsonResponseString = strings.TrimSpace(jsonResponseString)

	// --- STEP 7: RETURN RAW JSON ---

	return c.SendString(jsonResponseString)
}
