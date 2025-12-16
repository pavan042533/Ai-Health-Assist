package medication

type MappedMedication struct {
	DrugName      string  `json:"drug_name"`
	IngredientIDs []uint  `json:"ingredient_ids,omitempty"`
	Strength      string  `json:"strength"`
	Frequency     string  `json:"frequency"`
	Duration      string  `json:"duration"`
	MatchScore    float64 `json:"match_score"`
}

type MedInput struct {
	DrugName     string `json:"drug_name"`
	Strength     string `json:"strength"`
	Frequency    string `json:"frequency"`
	Duration     string `json:"duration"`
	Instructions string `json:"instructions"`
}

// Ingredient table
type DrugIngredient struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

// Product table
type DrugProduct struct {
	ID          uint   `gorm:"primaryKey"`
	ProductName string `gorm:"not null"`
	BrandName   string
	Strength    string
	Form        string

	// Product ↔ Ingredients (many-to-many)
	Ingredients []DrugIngredient `gorm:"many2many:product_ingredients;"`

	// Product ↔ Synonyms (1-to-many)
	Synonyms []DrugSynonym `gorm:"foreignKey:ProductID"`
}

// // Join table for many-to-many
// type ProductIngredient struct {
// 	ID            uint `gorm:"primaryKey"`
// 	DurgProductID uint
// 	DrugProduct   DrugProduct `gorm:"foreignKey:DurgProductID"`
// 	IngredientID  uint
// 	Ingredient    DrugIngredient `gorm:"foreignKey:"`
// }

// Synonyms table
type DrugSynonym struct {
	ID           uint   `gorm:"primaryKey"`
	Synonym      string `gorm:"index;not null"`
	ProductID    uint   `gorm:"index"`
	IngredientID uint   `gorm:"index"`
}

type BestMatch struct {
	Product DrugProduct
	Score   int
}
