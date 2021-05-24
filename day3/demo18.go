package day3

import "fmt"

func Demo18() {
	a := [3]string{"Hello", "world", "Hi"}
	b := [1]int{}
	c := [2]bool{0: true}
	d := &c

	fmt.Println(a, b, c, d)
	d[0] = false
	fmt.Println(c, d)
}
