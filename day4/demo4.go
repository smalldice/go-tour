package day4

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Demo4() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))
}
