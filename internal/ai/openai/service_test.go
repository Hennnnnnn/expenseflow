package openai

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestOpenAI(t *testing.T) {

	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY empty")
	}

	s := New(apiKey)

	result, err := s.Categorize(
		context.Background(),
		"STARBUCKS COFFEE",
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
