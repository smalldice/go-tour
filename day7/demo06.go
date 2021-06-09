package day7

import (
	"fmt"
	"math"
)

type Abser6 interface {
	Abs() float64
}

type Vertex6 struct {
	x, y float64
}

type MyFloat6 float64

func (v *Vertex6) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (f MyFloat6) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func Demo6() {
	var a Abser6

	f := MyFloat6(-math.Sqrt2)
	v := Vertex6{3, 4}
	a = f
	a = &v

	fmt.Println(a.Abs())
}
