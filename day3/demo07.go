package day3

import (
	"fmt"
	"math"
)

func myPow1(x, y, lim float64) float64 {
	if math.Pow(x, y) < lim {
		return math.Pow(x, y)
	}

	return lim
}

func Demo7() {
	fmt.Println(
		myPow1(2, 3, 10),
		myPow1(3, 3, 20),
	)
}
