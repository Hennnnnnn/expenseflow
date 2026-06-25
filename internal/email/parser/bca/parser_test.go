package bca

import "testing"

func TestNewParser(t *testing.T) {

	parser := New()

	if parser == nil {
		t.Fatal("parser should not be nil")
	}
}
