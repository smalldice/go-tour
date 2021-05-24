package day3

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Demo26() {
	// range pow 返回切片的 下标 和 下标对应的value
	for i, v := range pow {
		fmt.Printf("2 ** %d = %d \n", i, v)
	}

	// 可以用 _ 来忽略下标或者 值
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}
