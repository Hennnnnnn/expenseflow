package pipeline

import (
	"context"
	"log"
	"time"

	"github.com/Hennnnnnn/expenseflow/internal/ai"
)

type Service struct {
	categorizer ai.Categorizer
	provider    ai.Provider
}

func New(
	categorizer ai.Categorizer,
	provider ai.Provider,
) *Service {

	return &Service{
		categorizer: categorizer,
		provider:    provider,
	}
}

func (s *Service) Categorize(
	merchant string,
) ai.CategoryResult {

	// Rule Engine
	result := s.categorizer.Categorize(merchant)

	if result.Confidence >= 0.9 {
		return result
	}

	log.Println("Using OpenAI for:", merchant)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()

	aiResult, err := s.provider.Categorize(
		ctx,
		merchant,
	)

	if err != nil {
		log.Printf("OpenAI error: %v\n", err)

		// fallback ke Rule Engine
		return result
	}

	log.Printf(
		"OpenAI: %s -> %s (%.2f)\n",
		merchant,
		aiResult.Category,
		aiResult.Confidence,
	)

	return aiResult
}
