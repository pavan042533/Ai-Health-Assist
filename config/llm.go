package config

import (
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

var VisionLLM *ollama.LLM

func InitLLM() {
	llm, err := ollama.New(ollama.WithModel("gemma3"))
	if err != nil {
		// Do not exit the process if Ollama isn't running or the model isn't available.
		// Leave VisionLLM as nil and log a warning; handlers should check for nil.
		log.Printf("Warning: Failed to create Ollama client: %v", err)
		VisionLLM = nil
		return
	}
	VisionLLM = llm
	fmt.Println("Ollama LLM initialized:", VisionLLM)
}
