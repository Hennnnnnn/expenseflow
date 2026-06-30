package ai

import "context"

type Provider interface {
	Categorize(
		ctx context.Context,
		merchant string,
	) (CategoryResult, error)
}
