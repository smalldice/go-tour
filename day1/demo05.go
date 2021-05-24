package day1

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Demo5() {
	a, b := swap("hello", "world")
	fmt.Println("hello world swapped result is ", a, b)
}
