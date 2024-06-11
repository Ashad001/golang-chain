package main

import (
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

}

func run() error {
	_, err := openai.New()
	if err != nil {
		return err
	}
	return nil
}
