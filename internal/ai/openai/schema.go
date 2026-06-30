package openai

var CategorySchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"category": map[string]any{
			"type": "string",
		},
		"confidence": map[string]any{
			"type": "number",
		},
		"reason": map[string]any{
			"type": "string",
		},
	},
	"required": []string{
		"category",
		"confidence",
		"reason",
	},
	"additionalProperties": false,
}
