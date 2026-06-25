package parser

import (
	"testing"

	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
)

func TestCanParse(t *testing.T) {

	parser := NewBCAParser()

	msg := &imap.Message{
		From: "no-reply@bca.co.id",
	}

	if !parser.CanParse(msg) {
		t.Fatal("expected parser can parse BCA email")
	}
}

func TestParserSelection(t *testing.T) {

	p := NewBCAParser()

	msg := &imap.Message{
		From: "notification@bca.co.id",
	}

	if !p.CanParse(msg) {
		t.Fatal("parser should accept BCA sender")
	}
}
