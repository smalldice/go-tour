package day7

import "fmt"

type IPAddr [4]byte

func (IP IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", IP[0], IP[2], IP[2], IP[3])
}

func Demo10() {
	ip := IPAddr{192, 168, 0, 1}
	fmt.Println(ip)
}
