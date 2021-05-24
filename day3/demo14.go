// 指针

package day3

import "fmt"

var p1 *int

func printP() {
	fmt.Println(p)
}

func Demo14() {
	printP()
	i := 42

	p1 = &i

	fmt.Println(*p)
	*p1 = 21
	fmt.Println(i)
}
