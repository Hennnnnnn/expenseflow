package ai

type CategoryResult struct {
	Category   string
	Confidence float64
	Reason     string
}

// func (s *Service) Categorize(
// 	merchant string,
// ) CategoryResult {

// 	switch {

// 	case strings.Contains(strings.ToLower(merchant), "grab"):

// 		return CategoryResult{
// 			Category:   "Transportation",
// 			Confidence: 0.95,
// 		}

// 	case strings.Contains(strings.ToLower(merchant), "gopay"):

// 		return CategoryResult{
// 			Category:   "Transfer",
// 			Confidence: 0.95,
// 		}

// 	case strings.Contains(strings.ToLower(merchant), "mcd"):

// 		return CategoryResult{
// 			Category:   "Food",
// 			Confidence: 0.99,
// 		}

// 	case strings.Contains(strings.ToLower(merchant), "starbucks"):

// 		return CategoryResult{
// 			Category:   "Coffee",
// 			Confidence: 0.99,
// 		}

// 	default:

// 		return CategoryResult{
// 			Category:   "Uncategorized",
// 			Confidence: 0,
// 		}
// 	}
// }
