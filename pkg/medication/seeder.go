package medication

import (
	"fmt"

	"gorm.io/gorm"
)

func SeedDrugData(db *gorm.DB) {
	// avoid double seeding
	var count int64
	db.Model(&DrugIngredient{}).Count(&count)
	if count > 0 {
		return
	}

	fmt.Println("seeding started")

	// ----------------------------------------
	// 1. INGREDIENTS (19 total)
	// ----------------------------------------
	ingredients := []DrugIngredient{
		{Name: "Paracetamol"},       // 1
		{Name: "Amoxicillin"},       // 2
		{Name: "Ibuprofen"},         // 3
		{Name: "Azithromycin"},      // 4
		{Name: "Cetirizine"},        // 5
		{Name: "Levocetirizine"},    // 6
		{Name: "Montelukast"},       // 7
		{Name: "Clavulanic Acid"},   // 8
		{Name: "Chlorpheniramine"},  // 9
		{Name: "Phenylephrine"},     // 10
		{Name: "Pantoprazole"},      // 11
		{Name: "Vitamin D3"},        // 12
		{Name: "Calcium"},           // 13
		{Name: "Ambroxol"},          // 14
		{Name: "Levosalbutamol"},    // 15
		{Name: "Guaifenesin"},       // 16
		{Name: "Diclofenac"},        // 17
		{Name: "Cefixime"},          // 18
		{Name: "Omeprazole"},        // 19
	}
	db.Create(&ingredients)

	// ----------------------------------------
	// 2. PRODUCTS (17 total)
	// ----------------------------------------
	products := []DrugProduct{
		{ProductName: "Dolo 650", BrandName: "Dolo", Strength: "650 mg", Form: "tablet"}, // 1
		{ProductName: "Calpol 500", BrandName: "Calpol", Strength: "500 mg", Form: "tablet"}, // 2
		{ProductName: "Crocin 500", BrandName: "Crocin", Strength: "500 mg", Form: "tablet"}, // 3
		{ProductName: "Azee 500", BrandName: "Azee", Strength: "500 mg", Form: "tablet"}, // 4
		{ProductName: "Azithral 500", BrandName: "Azithral", Strength: "500 mg", Form: "tablet"}, // 5
		{ProductName: "Augmentin Duo 625", BrandName: "Augmentin", Strength: "625 mg", Form: "tablet"}, // 6
		{ProductName: "Sinarest Tablet", BrandName: "Sinarest", Strength: "Multiple", Form: "tablet"}, // 7
		{ProductName: "Pantocid 40", BrandName: "Pantocid", Strength: "40 mg", Form: "tablet"}, // 8
		{ProductName: "Pantoprazole 40 mg", BrandName: "Generic", Strength: "40 mg", Form: "tablet"}, // 9
		{ProductName: "Cetzine", BrandName: "Cetzine", Strength: "10 mg", Form: "tablet"}, // 10
		{ProductName: "Zyrtec", BrandName: "Zyrtec", Strength: "10 mg", Form: "tablet"}, // 11
		{ProductName: "Montair-LC", BrandName: "Montair-LC", Strength: "Multiple", Form: "tablet"}, // 12
		{ProductName: "Becosules Capsule", BrandName: "Becosules", Strength: "Multiple", Form: "capsule"}, // 13
		{ProductName: "Voveran 50", BrandName: "Voveran", Strength: "50 mg", Form: "tablet"}, // 14
		{ProductName: "Zifi 200", BrandName: "Zifi", Strength: "200 mg", Form: "tablet"}, // 15
		{ProductName: "Flexon", BrandName: "Flexon", Strength: "Multiple", Form: "tablet"}, // 16
		{ProductName: "Omez 20", BrandName: "Omez", Strength: "20 mg", Form: "capsule"}, // 17
	}
	db.Create(&products)

	
	// // ----------------------------------------
	// // 3. PRODUCT → INGREDIENT MAPPING
	// // ----------------------------------------
	db.Model(&products[0]).Association("Ingredients").Append(&DrugIngredient{ID: 1})  // Dolo 650 → Paracetamol
	db.Model(&products[1]).Association("Ingredients").Append(&DrugIngredient{ID: 1})  // Calpol 500 → Paracetamol
	db.Model(&products[2]).Association("Ingredients").Append(&DrugIngredient{ID: 1})	  // Crocin 500 → Paracetamol
	db.Model(&products[3]).Association("Ingredients").Append(&DrugIngredient{ID: 4})  // Azee 500 → Azithromycin
	db.Model(&products[4]).Association("Ingredients").Append(&DrugIngredient{ID: 4})  // Azithral 500 → Azithromycin
	db.Model(&products[5]).Association("Ingredients").Append(&DrugIngredient{ID: 2}, &DrugIngredient{ID: 8})  // Augmentin Duo 625 → Amoxicillin + Clavulanic Acid
	db.Model(&products[6]).Association("Ingredients").Append(&DrugIngredient{ID: 1}, &DrugIngredient{ID: 9}, &DrugIngredient{ID: 10})  // Sinarest → Paracetamol + Chlorpheniramine + Phenylephrine
	db.Model(&products[7]).Association("Ingredients").Append(&DrugIngredient{ID: 11})  // Pantocid 40 → Pantoprazole
	db.Model(&products[8]).Association("Ingredients").Append(&DrugIngredient{ID: 11})  // Pantoprazole 40 mg → Pantoprazole
	db.Model(&products[9]).Association("Ingredients").Append(&DrugIngredient{ID: 5})  // Cetzine → Cetirizine
	db.Model(&products[10]).Association("Ingredients").Append(&DrugIngredient{ID: 5})  // Zyrtec → Cetirizine
	db.Model(&products[11]).Association("Ingredients").Append(&DrugIngredient{ID: 7}, &DrugIngredient{ID: 6})  // Montair-LC → Montelukast + Levocetirizine
	db.Model(&products[12]).Association("Ingredients").Append(&DrugIngredient{ID: 12}, &DrugIngredient{ID: 13})  // Becosules Capsule → Vitamin D3 + Calcium
	db.Model(&products[13]).Association("Ingredients").Append(&DrugIngredient{ID: 17})  // Voveran 50 → Diclofenac
	db.Model(&products[14]).Association("Ingredients").Append(&DrugIngredient{ID: 18})  // Zifi 200 → Cefixime
	db.Model(&products[15]).Association("Ingredients").Append(&DrugIngredient{ID: 3}, &DrugIngredient{ID: 1})  // Flexon → Ibuprofen + Paracetamol
	db.Model(&products[16]).Association("Ingredients").Append(&DrugIngredient{ID: 19})  // Omez 20 → Omeprazole
	
	// mappings := []ProductIngredient{
	// 	// Paracetamol-based (Dolo, Calpol, Crocin)
	// 	{DurgProductID: 1, IngredientID: 1},
	// 	{DurgProductID: 2, IngredientID: 1},
	// 	{DurgProductID: 3, IngredientID: 1},

	// 	// Azithromycin products
	// 	{DurgProductID: 4, IngredientID: 4},
	// 	{DurgProductID: 5, IngredientID: 4},
	// 	// Augmentin = Amoxicillin + Clavulanic Acid
	// 	{DurgProductID: 6, IngredientID: 2},
	// 	{DurgProductID: 6, IngredientID: 8},

	// 	// Sinarest = Paracetamol + Chlorpheniramine + Phenylephrine
	// 	{DurgProductID: 7, IngredientID: 1},
	// 	{DurgProductID: 7, IngredientID: 9},
	// 	{DurgProductID: 7, IngredientID: 10},

	// 	// Pantoprazole products
	// 	{DurgProductID: 8, IngredientID: 11},
	// 	{DurgProductID: 9, IngredientID: 11},
	// 	// Cetirizine products
	// 	{DurgProductID: 10, IngredientID: 5},
	// 	{DurgProductID: 11, IngredientID: 5},

	// 	// Montair-LC = Montelukast + Levocetirizine
	// 	{DurgProductID: 12, IngredientID: 7},
	// 	{DurgProductID: 12, IngredientID: 6},
	// 	// Becosules = Vitamin D3 + Calcium
	// 	{DurgProductID: 13, IngredientID: 12},
	// 	{DurgProductID: 13, IngredientID: 13},

	// 	// Voveran
	// 	{DurgProductID: 14, IngredientID: 17},

	// 	// Zifi 200 = Cefixime
	// 	{DurgProductID: 15, IngredientID: 18},

	// 	// Flexon = Ibuprofen + Paracetamol
	// 	{DurgProductID: 16, IngredientID: 3},
	// 	{DurgProductID: 16, IngredientID: 1},
	// 	// Omez = Omeprazole
	// 	{DurgProductID: 17, IngredientID: 19},
	// }
	// db.Create(&mappings)

	// ----------------------------------------
	// 4. EXTENDED SYNONYMS (~150 entries)
	// ----------------------------------------
	synonyms := []DrugSynonym{

		// -----------------------------------
		// PARACETAMOL FAMILY
		// -----------------------------------
		{Synonym: "paracetamol", ProductID: 1, IngredientID: 1},
		{Synonym: "paracetmol", ProductID: 1, IngredientID: 1},
		{Synonym: "paracitamol", ProductID: 1, IngredientID: 1},
		{Synonym: "paracatamol", ProductID: 1, IngredientID: 1},
		{Synonym: "paracet", ProductID: 1, IngredientID: 1},
		{Synonym: "para", ProductID: 1, IngredientID: 1},
		{Synonym: "pcm", ProductID: 1, IngredientID: 1},
		{Synonym: "p-500", ProductID: 1, IngredientID: 1},
		{Synonym: "p500", ProductID: 1, IngredientID: 1},
		{Synonym: "paracetml", ProductID: 1, IngredientID: 1},
		{Synonym: "paracetmol", ProductID: 1, IngredientID: 1},

		// Dolo
		{Synonym: "dolo", ProductID: 1, IngredientID: 1},
		{Synonym: "dolo650", ProductID: 1, IngredientID: 1},
		{Synonym: "dolo 650", ProductID: 1, IngredientID: 1},
		{Synonym: "dol650", ProductID: 1, IngredientID: 1},
		{Synonym: "dollo", ProductID: 1, IngredientID: 1},
		{Synonym: "dloo", ProductID: 1, IngredientID: 1},
		{Synonym: "doloo", ProductID: 1, IngredientID: 1},

		// Calpol
		{Synonym: "calpol", ProductID: 2, IngredientID: 1},
		{Synonym: "calpol500", ProductID: 2, IngredientID: 1},
		{Synonym: "calpol 500", ProductID: 2, IngredientID: 1},
		{Synonym: "calpol tab", ProductID: 2, IngredientID: 1},

		// Crocin
		{Synonym: "crocin", ProductID: 3, IngredientID: 1},
		{Synonym: "crocin500", ProductID: 3, IngredientID: 1},
		{Synonym: "crocn", ProductID: 3, IngredientID: 1},
		{Synonym: "croci", ProductID: 3, IngredientID: 1},
		{Synonym: "crocn tab", ProductID: 3, IngredientID: 1},

		// -----------------------------------
		// AZITHROMYCIN FAMILY
		// -----------------------------------
		{Synonym: "azee", ProductID: 4, IngredientID: 4},
		{Synonym: "azee500", ProductID: 4, IngredientID: 4},
		{Synonym: "azee 500", ProductID: 4, IngredientID: 4},
		{Synonym: "aze 500", ProductID: 4, IngredientID: 4},
		{Synonym: "azitro", ProductID: 4, IngredientID: 4},
		{Synonym: "azithro", ProductID: 4, IngredientID: 4},
		{Synonym: "azth", ProductID: 4, IngredientID: 4},

		// Azithral
		{Synonym: "azithral", ProductID: 5, IngredientID: 4},
		{Synonym: "azithrl", ProductID: 5, IngredientID: 4},
		{Synonym: "azithral500", ProductID: 5, IngredientID: 4},
		{Synonym: "azthral", ProductID: 5, IngredientID: 4},
		{Synonym: "azithrol", ProductID: 5, IngredientID: 4},

		// -----------------------------------
		// AUGMENTIN (AMOX + CLAV)
		// -----------------------------------
		{Synonym: "augmentin", ProductID: 6, IngredientID: 2},
		{Synonym: "augmentin625", ProductID: 6, IngredientID: 2},
		{Synonym: "augmentin 625", ProductID: 6, IngredientID: 2},
		{Synonym: "augduo", ProductID: 6, IngredientID: 2},
		{Synonym: "augg", ProductID: 6, IngredientID: 2},
		{Synonym: "amoxicillin", ProductID: 6, IngredientID: 2},
		{Synonym: "amox", ProductID: 6, IngredientID: 2},
		{Synonym: "amoxy", ProductID: 6, IngredientID: 2},
		{Synonym: "amoxcillin", ProductID: 6, IngredientID: 2},
		{Synonym: "amoxycillin", ProductID: 6, IngredientID: 2},

		// -----------------------------------
		// SINAREST (Combo cold med)
		// -----------------------------------
		{Synonym: "sinarest", ProductID: 7, IngredientID: 1},
		{Synonym: "sina rest", ProductID: 7, IngredientID: 1},
		{Synonym: "sinarestt", ProductID: 7, IngredientID: 1},
		{Synonym: "sinrest", ProductID: 7, IngredientID: 1},
		{Synonym: "sinarst", ProductID: 7, IngredientID: 1},
		{Synonym: "synarest", ProductID: 7, IngredientID: 1},
		{Synonym: "snares", ProductID: 7, IngredientID: 1},
		{Synonym: "sinares", ProductID: 7, IngredientID: 1},
		{Synonym: "snarest", ProductID: 7, IngredientID: 1},

		// -----------------------------------
		// PANTOPRAZOLE FAMILY
		// -----------------------------------
		{Synonym: "pantoprazole", ProductID: 9, IngredientID: 11},
		{Synonym: "pantoprzole", ProductID: 9, IngredientID: 11},
		{Synonym: "pantopazole", ProductID: 9, IngredientID: 11},
		{Synonym: "panto", ProductID: 9, IngredientID: 11},
		{Synonym: "pantocid", ProductID: 8, IngredientID: 11},
		{Synonym: "pantacid", ProductID: 8, IngredientID: 11},
		{Synonym: "pantocid40", ProductID: 8, IngredientID: 11},

		// -----------------------------------
		// CETIRIZINE / LEVOCETIRIZINE
		// -----------------------------------
		{Synonym: "cetzine", ProductID: 10, IngredientID: 5},
		{Synonym: "cetizine", ProductID: 10, IngredientID: 5},
		{Synonym: "cetirizine", ProductID: 10, IngredientID: 5},
		{Synonym: "cetrizine", ProductID: 10, IngredientID: 5},
		{Synonym: "cetzin", ProductID: 10, IngredientID: 5},
		{Synonym: "zyrtec", ProductID: 11, IngredientID: 5},
		{Synonym: "zirtic", ProductID: 11, IngredientID: 5},
		{Synonym: "zirtac", ProductID: 11, IngredientID: 5},

		// -----------------------------------
		// MONTAIR-LC (Montelukast + Levocet)
		// -----------------------------------
		{Synonym: "montair", ProductID: 12, IngredientID: 7},
		{Synonym: "montair-lc", ProductID: 12, IngredientID: 7},
		{Synonym: "montlc", ProductID: 12, IngredientID: 7},

		// -----------------------------------
		// VOVERAN (Diclofenac)
		// -----------------------------------
		{Synonym: "voveran", ProductID: 14, IngredientID: 17},
		{Synonym: "voveran50", ProductID: 14, IngredientID: 17},
		{Synonym: "voveran 50", ProductID: 14, IngredientID: 17},
		{Synonym: "voveron", ProductID: 14, IngredientID: 17},
		{Synonym: "vovran", ProductID: 14, IngredientID: 17},

		// -----------------------------------
		// ZIFI (Cefixime)
		// -----------------------------------
		{Synonym: "zifi", ProductID: 15, IngredientID: 18},
		{Synonym: "zifi200", ProductID: 15, IngredientID: 18},
		{Synonym: "zifi 200", ProductID: 15, IngredientID: 18},
		{Synonym: "zifii", ProductID: 15, IngredientID: 18},
		{Synonym: "zify", ProductID: 15, IngredientID: 18},

		// -----------------------------------
		// FLEXON (Ibuprofen + Paracetamol)
		// -----------------------------------
		{Synonym: "flexon", ProductID: 16, IngredientID: 3},
		{Synonym: "flexton", ProductID: 16, IngredientID: 3},
		{Synonym: "flexn", ProductID: 16, IngredientID: 3},
		{Synonym: "flexon tab", ProductID: 16, IngredientID: 3},

		// -----------------------------------
		// OMEZ (Omeprazole)
		// -----------------------------------
		{Synonym: "omez", ProductID: 17, IngredientID: 19},
		{Synonym: "omez20", ProductID: 17, IngredientID: 19},
		{Synonym: "omez 20", ProductID: 17, IngredientID: 19},
		{Synonym: "omz", ProductID: 17, IngredientID: 19},
		{Synonym: "omeprazole", ProductID: 17, IngredientID: 19},
		{Synonym: "omprazole", ProductID: 17, IngredientID: 19},

	}

	db.Create(&synonyms)
}
