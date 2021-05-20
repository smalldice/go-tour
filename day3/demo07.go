package main

import (
	"fmt"
	"math"
)

func pow (x, y, lim float64) float64 {
	if math.Pow(x, y) < lim {
		return math.Pow(x, y)
	}

	return lim
}

func main () {
	fmt.Println(
		pow(2, 3, 10),
		pow(3, 3, 20),
	)
}