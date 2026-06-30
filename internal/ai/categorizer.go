package ai

type Categorizer interface {
	Categorize(merchant string) CategoryResult
}