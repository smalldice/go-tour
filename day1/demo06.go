package day1

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func Demo6() {
	fmt.Println(split(17))
}
