package day7

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	x, y float64
}

func (v *Vertex5) Scale(f float64) {
	v.x *= f
	v.y *= f
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo5() {
	v := &Vertex5{3, 4}

	v.Scale(10)
	fmt.Println(v.Abs())
}
