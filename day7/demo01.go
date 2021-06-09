package day7

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	x, y float64
}

func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo1() {
	v := Vertex1{2, 3}
	fmt.Println(v.Abs())
}
