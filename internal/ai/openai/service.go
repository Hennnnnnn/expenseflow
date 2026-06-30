package openai

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Hennnnnnn/expenseflow/internal/ai"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
)

type Service struct {
	client openai.Client
}

func New(apiKey string) *Service {

	log.Printf("API Key length: %d", len(apiKey))
	log.Printf("API Key prefix: %.10s", apiKey)

	return &Service{
		client: NewClient(apiKey),
	}
}

func (s *Service) Categorize(
	ctx context.Context,
	merchant string,
) (ai.CategoryResult, error) {

	resp, err := s.client.Responses.New(
		ctx,
		responses.ResponseNewParams{

			Model: openai.ChatModelGPT4_1Mini,

			Input: responses.ResponseNewParamsInputUnion{
				OfString: param.NewOpt(
					BuildPrompt(merchant),
				),
			},

			Text: responses.ResponseTextConfigParam{
				Format: responses.ResponseFormatTextConfigParamOfJSONSchema(
					"expense_category",
					CategorySchema,
				),
			},
		},
	)

	if err != nil {
		return ai.CategoryResult{}, err
	}

	var result AIResponse

	output := resp.OutputText()

	log.Println("========== RAW OUTPUT ==========")
	log.Println(output)
	log.Println("================================")

	output = strings.TrimSpace(output)
	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimPrefix(output, "```")
	output = strings.TrimSuffix(output, "```")
	output = strings.TrimSpace(output)

	err = json.Unmarshal(
		[]byte(output),
		&result,
	)

	if err != nil {
		return ai.CategoryResult{}, err
	}

	return ai.CategoryResult{
		Category:   result.Category,
		Confidence: result.Confidence,
		Reason:     result.Reason,
	}, nil
}

var _ ai.Provider = (*Service)(nil)
