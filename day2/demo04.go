package day2

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe1   bool       = false
	MaxInt1 uint64     = 1<<64 - 1
	z1      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Demo4() {
	fmt.Printf("Type: %T Value: %v\n", ToBe1, ToBe1)
	fmt.Printf("Type: %T Value: %v\n", MaxInt1, MaxInt1)
	fmt.Printf("Type: %T Value: %v\n", z1, z1)
}
