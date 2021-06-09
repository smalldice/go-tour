package day7

import (
	"fmt"
	"math"
	"time"
)

type SqrtError struct {
	When time.Time
	What string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &SqrtError{
			time.Now(),
			fmt.Sprintf("%v is an negative number", x),
		}
	}

	return math.Sqrt(x), nil
}

func Demo12() {
	x := float64(-2)
	r, e := Sqrt(x)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(r)
}
