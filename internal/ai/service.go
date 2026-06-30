package ai

import "strings"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Categorize(
	merchant string,
) CategoryResult {

	merchant = strings.ToLower(merchant)

	for keyword, category := range MerchantRules {

		if strings.Contains(merchant, keyword) {

			return CategoryResult{
				Category:   category,
				Confidence: 0.95,
			}
		}
	}

	return CategoryResult{
		Category:   "Uncategorized",
		Confidence: 0,
	}
}

var _ Categorizer = (*Service)(nil)
