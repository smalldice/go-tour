package day5

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func (p Person) howOld() int {
	fmt.Printf("How old , address is %p\n", &p)
	return p.age
}

func (p *Person) howOld2() int {
	fmt.Printf("How old, address is %p\n", p)
	return p.age
}

func (p Person) who() string {
	fmt.Printf("What is the person's name, address is %p\n", &p)
	return p.name
}

func (p *Person) who2() string {
	fmt.Printf("What is the person's name, address is %p \n", p)
	return p.name
}

func (p Person) growUp() int {
	p.age += 1
	return p.age
}

func (p *Person) growUp2() int {
	p.age += 1
	return p.age
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Demo5() {
	xjh := Person{"xjh", 18}
	xjh.howOld()
	xjh.who()
	xjh.growUp()
	fmt.Println(xjh.age)
	xjh.howOld2()
	xjh.who2()
	xjh.growUp2()
	fmt.Println(xjh.age)

	v := Vertex1{3, 4}
	fmt.Println(v.Abs())
}
