package day5

import "fmt"

type I8 interface {
	M8()
}

type T8 struct {
	S string
}

func (t *T8) M8() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func Demo8() {
	var i I8
	var t *T8

	i = t
	describe8(i)
	i.M8()

	i = &T8{"Hello"}
	describe8(i)
	i.M8()

}

func describe8(i I8) {
	fmt.Printf("(%v, %T)", i, i)
}
