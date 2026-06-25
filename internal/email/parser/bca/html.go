package bca

import (
	"io"

	"github.com/emersion/go-message"
	"github.com/emersion/go-message/mail"
)

type HTMLExtractor struct {
}

func NewHTMLExtractor() *HTMLExtractor {
	return &HTMLExtractor{}
}

func (e *HTMLExtractor) Extract(entity *message.Entity) (string, error) {

	body, err := io.ReadAll(entity.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (e *HTMLExtractor) ExtractHTML(entity *message.Entity) (string, error) {

	mr := mail.NewReader(entity)

	for {
		part, err := mr.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return "", err
		}

		switch h := part.Header.(type) {

		case *mail.InlineHeader:

			contentType, _, _ := h.ContentType()

			if contentType == "text/html" {

				body, err := io.ReadAll(part.Body)

				if err != nil {
					return "", err
				}

				return string(body), nil
			}
		}
	}

	return "", nil
}
