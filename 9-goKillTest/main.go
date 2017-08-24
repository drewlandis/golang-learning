package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Process started...")

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)

	go func() {
		<-term
		fmt.Println("Terminating...")
		sleep := 60
		for i := 0; i <= sleep; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Goodbye...")
		os.Exit(0)
	}()

	time.Sleep(5 * time.Minute)
}
