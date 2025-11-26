package main

import (
	"ai_health_assistant/config"
	"log"

	"github.com/gofiber/fiber/v2"

	// "ai_health_assistant/pkg"
	"ai_health_assistant/public"
)

func main() {
	config.Initconfig()
	// Initialize the vision LLM so handlers can use `config.VisionLLM`.
	// This will log a warning and continue if the Ollama client cannot be created.
	config.InitLLM()
	// pkg.Init()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	public.MountRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
