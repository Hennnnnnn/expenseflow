package bca

import (
	"bytes"

	"golang.org/x/net/html"
)

type TextExtractor struct {
}

func NewTextExtractor() *TextExtractor {
	return &TextExtractor{}
}

func (e *TextExtractor) Extract(root *html.Node) string {

	var buf bytes.Buffer

	var walk func(*html.Node)

	walk = func(n *html.Node) {

		if n.Type == html.TextNode {

			buf.WriteString(n.Data)
			buf.WriteString("\n")

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(root)

	return buf.String()
}
