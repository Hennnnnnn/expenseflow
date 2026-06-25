package bca

import "os"

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(path string) ([]byte, error) {

	return os.ReadFile(path)

}
