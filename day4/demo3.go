package day4

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	parts := strings.Fields(s)
	m := make(map[string]int)

	for i, v := range parts {
		fmt.Println(i, v)
		if m[v] != 0 {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	fmt.Println(m)
	return m
}

func Demo3() {
	wc.Test(WordCount)
}
