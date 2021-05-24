package day3

import "fmt"

func Demo19() {
	primes := [6]int{1, 2, 3, 4}

	var s []int = primes[1:4]

	fmt.Println(s)
}
