package day5

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	x, y float64
}

func (v *Vertex4) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo4() {
	v := &Vertex4{3, 4}

	fmt.Printf("Before Scaling : %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After Scaling: %+v, Abs: %v\n", v, v.Abs())
}
