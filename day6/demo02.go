package day6

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}
type ErrEmptyBuffer []byte

func (b ErrEmptyBuffer) Error() string {
	return fmt.Sprintf("cannot read an empty buffer: %b \n", b)
}

func (r MyReader) Read(b []byte) (int, error) {
	bLength := len(b)

	if bLength == 0 {
		return 0, ErrEmptyBuffer(b)
	}

	for i := range b {
		b[i] = 'A'
	}

	return bLength, nil
}

func Demo2() {
	r := MyReader{}
	b := make([]byte, 4)
	len, err := r.Read(b)
	println(b, len, err)
	reader.Validate(MyReader{})
}
