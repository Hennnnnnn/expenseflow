package parser

import "strings"

func Normalize(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	text = strings.TrimSpace(text)
	return text
}