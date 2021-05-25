package day5

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number %v", float64(e))
}

const EPSILON = 0.00001

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := x / 2
		i := 1

		for math.Abs(z*z-x) > EPSILON {
			fmt.Println(math.Abs(z*z - x))
			i += 1
			z -= (z*z - x) / (2 * z)
		}

		fmt.Println(i)

		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func Demo14() {
	// Sqrt(199)
	x := -2
	result, e := Sqrt(-2)

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("%v sqrt value is %v \n", x, result)
}
