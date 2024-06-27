package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func InvokeChain(apiKey string, userQuery string) (ChatOutput, error) {
	log.Println(userQuery)
	llm, err := openai.New(
		openai.WithModel("llama3-8b-8192"),
		openai.WithBaseURL("https://api.groq.com/openai/v1"),
		openai.WithToken(apiKey),
	)
	if err != nil {
		return ChatOutput{}, fmt.Errorf("error creatnig OpenAI client: %v", err)
	}

	systemPrompt, err := os.ReadFile("./prompts/system_prompt.txt")
	if err != nil {
		return ChatOutput{}, fmt.Errorf("error reading system prompt file: %v", err)
	}

	prompt := string(systemPrompt) + "\n\nUser Query: " + userQuery

	// Write prompt to log file ==> To be removed in production
	err = writePromptToLogFile(prompt)
	if err != nil {
		return ChatOutput{}, fmt.Errorf("error writing prompt to log file: %v", err)
	}

	ctx := context.Background()

	var responseContent string
	responseContent, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
		llms.WithTemperature(0.8),
		llms.WithJSONMode(),
	)

	if err != nil {
		return ChatOutput{}, fmt.Errorf("error in generating content: %v", err)
	}
	log.Printf("Raw response: %s", responseContent)

	responseContent = extractResponse(responseContent)

	var chatOutput ChatOutput
	if err := json.Unmarshal([]byte(responseContent), &chatOutput); err != nil {
		return ChatOutput{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return chatOutput, nil
}

func extractResponse(responseContent string) string {
	re := regexp.MustCompile("```(.*?)```")
	matches := re.FindStringSubmatch(responseContent)
	if len(matches) > 1 {
		return matches[1]
	}
	return responseContent
}

func writePromptToLogFile(prompt string) error {
	// if _, err := os.Stat("./log"); os.IsNotExist(err) {
	// 	os.Mkdir("./log", 0755)
	// }
	file, err := os.OpenFile("./prompt.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(prompt); err != nil {
		return fmt.Errorf("error writing to log file: %v", err)
	}

	return nil
}
