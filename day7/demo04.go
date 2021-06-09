package day7

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	x, y float64
}

func Scale(v *Vertex4, f float64) {
	v.x *= f
	v.y *= f
}

func Abs(v Vertex4) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo4() {
	v := Vertex4{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
