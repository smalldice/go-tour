package day5

import "fmt"

func Demo9() {
	var i interface{} = "Hello"

	s := i.(string)

	fmt.Println(s)

	s, ok := i.(string)

	fmt.Println(s, ok)

	// f, ok := i.(float64)

	// fmt.Println(f, ok)

	f, ok := i.(float64)

	fmt.Println(f, ok)
}
