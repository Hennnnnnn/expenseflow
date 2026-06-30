package openai

import "fmt"

func BuildPrompt(merchant string) string {

	return fmt.Sprintf(`
You are a finance categorization engine.

Categorize this merchant.

Merchant:
%s

Return ONLY JSON.

{
  "category":"Food",
  "confidence":0.97,
  "reason":"Coffee shop"
}
`, merchant)

}
