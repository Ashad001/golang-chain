package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	content, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(content)
}

func run() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY not set in .env file")
	}

	llm, err := openai.New()
	if err != nil {
		return "", fmt.Errorf("error creating OpenAI client: %v", err)
	}

	fmt.Println("OpenAI client created successfully")

	ctx := context.Background()
	prompt := "What is GoLang?"

	content, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
	)

	if err != nil {
		return "", fmt.Errorf("error in generating content: %v", err)
	}

	fmt.Println("Generated content:", content)

	return content, nil
}
