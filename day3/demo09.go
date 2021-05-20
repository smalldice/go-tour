package main

import (
	"fmt"
	"math"
)

const EPSILON = 0.00001


func Sqrt(x float64) float64 {
	z  := 1.0
	i  := 1


	for math.Abs(z * z - x) > EPSILON {
		fmt.Println(math.Abs(z * z - x))
		i += 1
		z -= (z * z - x) / (2 * z)
	}

	fmt.Println(i)

	return z
}

func main() {
	fmt.Println("平方根：", Sqrt(10))
	fmt.Println("平方根：", math.Sqrt(10))
}
