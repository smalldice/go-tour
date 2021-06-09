package day7

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func Demo7() {
	var i I = T{"Hello"}
	i.M()
}
