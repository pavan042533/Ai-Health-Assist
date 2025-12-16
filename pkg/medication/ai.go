package medication

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

// AskAIForIngredient uses Gemini API to get generic ingredient name(s) for a given brand drug name
func AskAIForIngredient(drugName string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", errors.New("no GEMINI_API_KEY set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")

	prompt := `
You are a medical assistant. Given a brand-name medicine, return ONLY the generic active ingredient name(s) as a comma-separated list, no extra text or explanation.

Examples:
"Dolo 650" -> "Paracetamol"
"Sinarest" -> "Paracetamol, Chlorpheniramine, Phenylephrine"
"Azee 500" -> "Azithromycin"

Now answer:
` + drugName + ` ->`

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if len(resp.Candidates) == 0 {
		return "", nil
	}

	// collect raw text
	result := ""
	for _, part := range resp.Candidates[0].Content.Parts {
		if t, ok := part.(genai.Text); ok {
			result += string(t)
		}
	}

	result = strings.TrimSpace(result)
	// remove trailing punctuation and newlines
	result = strings.ReplaceAll(result, "\n", " ")
	result = strings.TrimSuffix(result, ".")
	result = strings.TrimSpace(result)

	return result, nil
}

// Helper: ensure ingredient exists, return its ID
func ensureIngredient(db *gorm.DB, name string) (uint, error) {
	name = strings.TrimSpace(name)
	lname := strings.ToLower(name)
	var ing DrugIngredient
	if err := db.Where("LOWER(name) = ?", lname).First(&ing).Error; err == nil {
		return ing.ID, nil
	}
	// create
	ing = DrugIngredient{Name: name}
	if err := db.Create(&ing).Error; err != nil {
		return 0, err
	}
	return ing.ID, nil
}

// Helper: create product + mapping + synonym when AI suggests ingredient(s)
func createProductFromAI(db *gorm.DB, brandName string, ingredientNames []string) (DrugProduct, error) {
	brandTrim := strings.TrimSpace(brandName)
	product := DrugProduct{
		ProductName: brandTrim,
		BrandName:   brandTrim,
		Strength:    "",
		Form:        "",
	}
	// create product
	if err := db.Create(&product).Error; err != nil {
		return DrugProduct{}, err
	}

	// for each ingredient ensure ingredient row and create mapping
	for _, in := range ingredientNames {
		in = strings.TrimSpace(in)
		if in == "" {
			continue
		}
		ingID, err := ensureIngredient(db, in)
		if err != nil {
			return DrugProduct{}, err
		}
		// m := ProductIngredient{
		// 	DurgProductID: product.ID,
		// 	IngredientID: ingID,
		// }
		// if err := db.Create(&m).Error; err != nil {
		// 	return DrugProduct{}, err
		// }
		// create mapping
		if err := db.Model(&product).Association("Ingredients").Append(&DrugIngredient{ID: ingID}); err != nil {
			return DrugProduct{}, err
		}
		// create a synonym entry for brand → ingredient (so future fuzzy can match)
		syn := DrugSynonym{
			Synonym:      strings.ToLower(brandTrim),
			ProductID:    product.ID,
			IngredientID: ingID,
		}
		if err := db.Create(&syn).Error; err != nil {
			return DrugProduct{}, err
		}
	}
	// reload with preloads
	_ = db.Preload("Ingredients").Preload("Synonyms").First(&product, product.ID).Error
	return product, nil
}
