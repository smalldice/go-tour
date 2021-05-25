package day5

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func Demo6() {
	var i I = &T{"Hello"}

	i.M()
}
