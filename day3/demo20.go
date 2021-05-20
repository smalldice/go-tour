package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"George",
		"Paul",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:2]

	fmt.Println(a, b)

	b[0] = "Hello"

	fmt.Println(names, a, b)
}
