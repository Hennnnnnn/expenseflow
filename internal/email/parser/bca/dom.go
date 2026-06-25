package bca

import (
	"strings"

	"golang.org/x/net/html"
)

type DOMParser struct {
}

func NewDOMParser() *DOMParser {
	return &DOMParser{}
}

func (d *DOMParser) Parse(htmlString string) (*html.Node, error) {

	return html.Parse(strings.NewReader(htmlString))

}
