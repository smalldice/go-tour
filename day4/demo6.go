package day4

import "fmt"

func fibonacci() func() int {
	n := []int{0, 1}
	i := 1

	return func() int {
		var num int
		switch i {
		case 1:
			num = n[0]
		case 2:
			num = n[1]
		default:
			numLen := len(n)
			num = n[numLen-1] + n[numLen-2]
			n = append(n, num)
		}
		i += 1

		return num
	}
}

func Demo6() {
	f := fibonacci()

	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
