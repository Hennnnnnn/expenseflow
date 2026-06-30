package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	key := os.Getenv("OPENAI_API_KEY")

	fmt.Println("Key length:", len(key))

	client := openai.NewClient(
		option.WithAPIKey(key),
	)

	models, err := client.Models.List(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(models)
}
