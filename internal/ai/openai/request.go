package openai

import (
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
)

type Request struct {
	Merchant string
}

func BuildRequest(
	merchant string,
) responses.ResponseNewParams {

	return responses.ResponseNewParams{

		Model: openai.ChatModelGPT4_1Mini,

		Input: responses.ResponseNewParamsInputUnion{
			OfString: param.NewOpt(
				BuildPrompt(merchant),
			),
		},

		Temperature: param.NewOpt(0.2),

		Text: responses.ResponseTextConfigParam{

			Format: responses.ResponseFormatTextConfigParamOfJSONSchema(
				"expense_category",
				CategorySchema,
			),
		},
	}
}
