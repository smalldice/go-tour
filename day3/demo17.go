package main

import "fmt"

func main() {
	var a [2]string // [ ]
	var b [6]int    //
	fmt.Println(a)
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	premies := [6]int{2, 3, 4, 6, 7, 11}

	for i := 0; i < len(premies); i++ {
		fmt.Println(premies[i])
		b[i] = premies[i] * 2
	}

	fmt.Println(b)
}
