package openai

import (
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func NewClient(apiKey string) openai.Client {
	return openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}
