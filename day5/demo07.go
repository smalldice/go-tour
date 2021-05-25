package day5

import (
	"fmt"
	"math"
)

type I7 interface {
	M7()
}

type T7 struct {
	S string
}

func (t *T) M7() {
	fmt.Println(t.S)
}

type F7 float64

func (f F7) M7() {
	fmt.Println(f)
}

func describe(i I7) {
	fmt.Printf("%v, %T , %p\n", i, i, &i)
}

func Demo7() {
	var i I7
	i = &T{"hello"}
	fmt.Printf("Address: %p\n", i)
	describe(i)
	i.M7()

	i = F7(math.Pi)
	fmt.Printf("Address: %p\n", &i)
	describe(i)
	i.M7()
}
