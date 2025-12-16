package config

import (
	"ai_health_assistant/models"
	"log"

	"fmt"
	"os"

	"ai_health_assistant/pkg/medication"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Cfg = GetConfig()

func InitDB() {
	log.Println("DB DSN:", os.Getenv("DB_HOST")) // Add this line
	log.Printf("Config: %+v\n", Cfg)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", Cfg.DbHost, Cfg.DbUser, Cfg.DbPassword, Cfg.DbName, Cfg.DbPort)
	log.Println("Connecting to DB...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	fmt.Println("Connected to DB")
	db = db.Debug()
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Medication{}, &medication.DrugIngredient{}, &medication.DrugProduct{}, &medication.DrugSynonym{})

	isSeedDB := os.Getenv("SEED_DB")
	if isSeedDB == "true" {
		log.Println("Seeding database with initial medication data...")
		medication.SeedDrugData(DB)
	}
}
