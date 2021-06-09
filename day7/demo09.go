package day7

import "fmt"

type I9 interface {
	M()
}

type T9 struct {
	S string
}

func (t *T9) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func describe9(i I) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func Demo9() {
	var i I
	var t *T
	i = t
	describe9(i)
	i.M()
}
