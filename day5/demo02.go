package day5

import "fmt"

type Vertex2 struct {
	x, y float64
}

func (v *Vertex2) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex2, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Demo2() {
	v := Vertex2{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 5)

	p := &Vertex2{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
