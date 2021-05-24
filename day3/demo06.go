package day3

import (
	"fmt"
	"math"
)

func myPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func Demo6() {
	fmt.Println(
		myPow(3, 2, 10),
		myPow(3, 3, 20),
	)
}
