package day7

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	x, y float64
}

func (v *Vertex3) Scale(s float64) {
	v.x *= s
	v.y *= s
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo3() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
