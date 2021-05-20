package main

import "fmt"

func main() {
	i := 4
	f := float64(i)
	u := uint(f)

	fmt.Println(i, f, u)
}