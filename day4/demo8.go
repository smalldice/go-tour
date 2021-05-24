package day4

import (
	"fmt"
	"math"
)

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func Demo8() {
	f := myFloat(-math.Sqrt2)

	fmt.Println(f.Abs())
}
