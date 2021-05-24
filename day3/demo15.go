package day3

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

type mat2 struct {
	x int
	y int
}

func Demo15() {
	v1 := Vertex1{1, 2}
	fmt.Println(v1.X, v1.Y)
	v1.X = 4
	fmt.Println(v1.X)

	v2 := v1
	fmt.Println(v2.X, v2.Y)
	v2.X = 10
	fmt.Println(v2.X, v1.X)

	v3 := &v2
	fmt.Println(v3.X, v3.Y)
	v3.X = 1
	fmt.Println(v3.X, v2.X)

	m1 := mat2{0, 1}
	m2 := m1

	fmt.Println(m1.x, m1.y)
	m2.x = 2
	fmt.Println(m2.x, m1.x)
}
