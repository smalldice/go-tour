package day5

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	x, y float64
}

func Abs(v Vertex1) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v *Vertex1, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Demo1() {
	v := Vertex1{3, 4}

	Scale(&v, 10)
	fmt.Println(Abs(v))
}
