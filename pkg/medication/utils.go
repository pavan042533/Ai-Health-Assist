package medication

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"gorm.io/gorm"
)

// Normalize string
func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

const FuzzyTrustThreshold = 0.70

func FindBestDrugMatchWithAI(db *gorm.DB, input string) (DrugProduct, float64, error) {
	// 1️ Try fuzzy match first
	product, score, err := FindBestDrugMatch(db, input)
	if err != nil {
		return DrugProduct{}, 0.0, err
	}
	if score > FuzzyTrustThreshold {
		return product, score, nil
	}

	// 2️ No good match found → AI fallback
	ingredientName, err := AskAIForIngredient(input)
	if err != nil || ingredientName == "" {
		// Total failure
		return DrugProduct{}, 0, err
	}

	// 3. AI fallback
	aiResponse, err := AskAIForIngredient(input)
	if err != nil {
		return DrugProduct{}, 0.0, err
	}
	if strings.TrimSpace(aiResponse) == "" {
		return DrugProduct{}, 0.0, nil
	}

	// AI returns comma-separated ingredient names
	parts := strings.Split(aiResponse, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	// create product + ingredient mappings + synonyms
	createdProduct, err := createProductFromAI(db, input, parts)
	if err != nil {
		return DrugProduct{}, 0.0, err
	}

	return createdProduct, 1.0, nil
}

// Main matching function
func FindBestDrugMatch(db *gorm.DB, input string) (DrugProduct, float64, error) {
	inputNorm := normalize(input)

	// quick exact match on synonym (case-insensitive)
	var exact DrugSynonym
	if err := db.Where("LOWER(synonym) = ?", inputNorm).First(&exact).Error; err == nil {
		// load product and ingredients
		var p DrugProduct
		if err := db.Preload("Ingredients").Preload("Synonyms").First(&p, exact.ProductID).Error; err == nil {
			return p, 1.0, nil
		}
	}

	//  prefix / LIKE based  fetch to reduce search space
	var fussySynonyms []DrugSynonym
	pattern := "%" + inputNorm + "%"
	err := db.Where("LOWER(synonym) LIKE ? OR LOWER(synonym) LIKE ?", inputNorm+"%", pattern).Limit(200).Find(&fussySynonyms).Error
	if err != nil {
		return DrugProduct{}, 0.0, err
	}

		if len(fussySynonyms) == 0 {
			// nothing even approximate matched
		return DrugProduct{}, 0.0, nil
	}

	// fuzzy match among candidates
	bestScore := -1
	var bestSyn DrugSynonym
	for _, s := range fussySynonyms {
		score := fuzzy.RankMatch(normalize(s.Synonym), inputNorm)
		if score > bestScore {
			bestScore = score
			bestSyn = s
		}
	}
	if bestScore < 30 {
		// too low to trust
		return DrugProduct{}, 0.0, nil
	}
	// load best product
	var product DrugProduct
	if err := db.Preload("Ingredients").Preload("Synonyms").First(&product, bestSyn.ProductID).Error; err != nil {
		return DrugProduct{}, 0.0, err
	}
	return product, float64(bestScore) / 100.0, nil
}
