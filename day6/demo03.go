package day6

import (
	"fmt"
	"time"
)

func Demo3() {
	tm := time.Unix(1622199261709/1000, 0)
	fmt.Println(tm.Format("2006-01-02 15:04"))
}
