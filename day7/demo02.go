package day7

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func Demo2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
