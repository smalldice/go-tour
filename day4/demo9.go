package day4

import "fmt"

type Vertex9 struct {
	x, y float64
}

func (v Vertex9) Abs9() float64 {
	return v.x + v.y
}

func (v Vertex9) Scale9(s float64) {
	v.x = v.x * s
	v.y = v.y * s
}

func Demo9() {
	v := Vertex9{1, 2}
	v.Scale9(10)

	fmt.Println(v.Abs9())
}
