package main

// This program will print "ping" forever. Hit enter to stop it.
/*
 * When pinger attempts to send a message on the channel, it will wait
 * until printer is ready to receive the message. (This is blocking)
 */

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) { // send only channel
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) { // bi-directional channel
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { // receive only channel
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
