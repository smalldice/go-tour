package day5

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	x, y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func AbsFunc(v *Vertex3) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo3() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(&v))

	p := &Vertex3{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(p))
}
