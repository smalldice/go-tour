package day4

import "fmt"

func Demo2() {
	m := make(map[string]int)

	m["Answer"] = 42

	fmt.Println(m["Answer"])

	m["Answer"] = 48

	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m["Answer"])

	v, ok := m["Answer"]

	fmt.Println(v, ok)

}
