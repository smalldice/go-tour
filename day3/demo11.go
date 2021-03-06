package day3

import (
	"fmt"
	"time"
)

func Demo11() {
	fmt.Print("When's Saturday? ")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("The day after tomorrow")
	default:
		fmt.Println("Too far away")
	}
}
