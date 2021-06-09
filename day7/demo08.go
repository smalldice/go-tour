package day7

import "fmt"

type I7 interface {
	M()
}

type T7 struct {
	S string
}

type MyF8 float64

func (t MyF8) M() {
	fmt.Println(t)
}

func (t T7) M() {
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}

func Demo8() {
	var i I

	i = &T7{"Hello"}
	describe(i)
	i.M()

	i = MyF8(2)
	describe(i)
}
