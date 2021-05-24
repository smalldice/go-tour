package day3

import "fmt"

func Demo22() {
	s := []int{2, 3, 5, 7, 11, 13}

	s1 := s[:0]
	printSlice1(s1)
	s2 := s[:4]
	printSlice1(s2)
	s3 := s[:2]
	printSlice1(s3)

	var s5 []string
	if s5 == nil {
		fmt.Println("nil!", len(s5), cap(s5))
	}
}

func printSlice1(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
