package day3

import "fmt"

func Demo12() {
	defer fmt.Println("world")

	fmt.Println("hello ")
}
