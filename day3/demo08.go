package day3

import (
	"fmt"
	"math"
)

func myPow2(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}

	return lim
}

func Demo8() {
	fmt.Println(
		myPow2(3, 2, 10),
		myPow2(3, 3, 20),
	)
}
