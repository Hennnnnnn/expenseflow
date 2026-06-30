package openai

import (
	"encoding/json"

	"github.com/Hennnnnnn/expenseflow/internal/ai"
)

func (s *Service) parseResult(
	output string,
) (ai.CategoryResult, error) {

	var result AIResponse

	err := json.Unmarshal(
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
