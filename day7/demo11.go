package day7

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"Oops! some error just happended",
	}
}

func Demo11() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
