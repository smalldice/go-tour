package day4

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func (v Vertex7) Abs() float64 {
	return math.Sqrt(v.X * v.Y)
}

var VertexInstace7 = Vertex7{3, 4}

func Demo7(v Vertex7) {
	fmt.Println(v.Abs())
}
