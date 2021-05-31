package day6

import (
	"fmt"
	"io"
	"strings"
)

func Demo1() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q \n", b[:n])

		if err == io.EOF {
			break
		}
	}
	fmt.Println(r, b)
}
