// 指针

package main

import "fmt"

var p *int

func printP () {
	fmt.Println(p)
}

func main () {
	printP()
	i := 42

	p = &i

	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}