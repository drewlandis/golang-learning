package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hour := time.Now().Hour()
	greeting, err := getGreeting(hour)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 8 {
		err := errors.New("Too early for greetings!")
		return message, err
	} else if hour < 12 {
		message = "Good Morning!"
	} else if hour < 18 {
		message = "Good Afternoon!"
	} else {
		message = "Good evening!"
	}

	return message, nil
}
