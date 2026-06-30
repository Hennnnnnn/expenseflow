package ai

import "testing"

func TestCategorizeGrab(t *testing.T) {
	service := New()

	result := service.Categorize("Grab* A-123")

	if result.Category != "Transportation" {
		t.Fatalf("expected Transportation, got %s", result.Category)
	}
}

func TestCategorize(t *testing.T) {
	tests := []struct {
		name     string
		merchant string
		expected string
	}{
		{"Grab", "Grab* A-123", "Transportation"},
		{"GoPay", "GoPayID", "Transfer"},
		{"Starbucks", "SBUX DT KM 88B", "Coffee"},
		{"McD", "MCDONALD'S 21901", "Food"},
		{"Claude", "ANTHROPIC* CLAUDE SUB", "AI Subscription"},
		{"Unknown", "ABC XYZ", "Uncategorized"},
	}

	service := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Categorize(tt.merchant)
			if result.Category != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, result.Category)
			}
		})
	}
}
