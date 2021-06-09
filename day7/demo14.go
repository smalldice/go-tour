package day7

import (
	"fmt"
	"io"
	"strings"
)

func Demo14() {
	r := strings.NewReader("HELLO READER")

	b := make([]byte, 8)
	fmt.Println(b)

	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v, err = %v, b = %v\n", n, b, err)

		if err == io.EOF {
			break
		}
	}
}
