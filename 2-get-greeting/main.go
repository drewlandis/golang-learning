package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	fmt.Println(getGreeting(hour))
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning!"
	} else if hour < 18 {
		return "Good Afternoon!"
	}
	return "Good Evening!"
}
