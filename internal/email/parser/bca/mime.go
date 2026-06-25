package bca

import (
	"bytes"

	"github.com/emersion/go-message"
)

type MIMEReader struct {
}

func NewMIMEReader() *MIMEReader {
	return &MIMEReader{}
}

func (r *MIMEReader) Parse(data []byte) (*message.Entity, error) {
	return message.Read(bytes.NewReader(data))
}
