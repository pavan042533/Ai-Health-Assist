package config

import (
	"ai_health_assistant/models"
	"log"

	"fmt"
	"os"

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
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Prescription{}, &models.ScannedMedication{})
}

