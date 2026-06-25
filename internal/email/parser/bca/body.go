package bca

import (
	"io"

	"github.com/emersion/go-message"
)

type BodyReader struct {
}

func NewBodyReader() *BodyReader {
	return &BodyReader{}
}

func (r *BodyReader) Read(entity *message.Entity) ([]byte, error) {

	return io.ReadAll(entity.Body)

}
