package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s1 := s[:0]
	printSlice(s1)
	s2 := s[:4]
	printSlice(s2)
	s3 := s[:2]
	printSlice(s3)

	var s5 []string
	if s5 == nil {
		fmt.Println("nil!", len(s5), cap(s5))
	}
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
