package main

import "fmt"

func main() {
	loop1()
	loop2()
	loopWithBreak()
}

func loop1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loop2() {
	i := 5
	lessThanTen := true
	for lessThanTen {
		if i >= 10 {
			lessThanTen = false
		}
		fmt.Println(i)
		i++
	}
}

func loopWithBreak() {
	i := 10
	for {
		if i >= 15 {
			break
		}
		fmt.Println(i)
		i++
	}
}
