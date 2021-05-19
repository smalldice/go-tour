package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main () {
	fmt.Print("1 plus 3 value is ", add(1, 3), "\n")
	fmt.Print("1 subtract 3 value is ", sub(1, 3), "\n")
}