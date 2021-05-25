package day5

import "fmt"

type Person11 struct {
	name string
	age  int
}

func (p Person11) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func Demo11() {
	a := Person11{"Arthur Dent", 42}
	z := Person11{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
