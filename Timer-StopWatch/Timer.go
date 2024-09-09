package main

import (
	"fmt"
	"time"
)

func startTimer(duration int) {
	for duration > 0 {
		fmt.Printf("\rTime remaining: %d seconds", duration)
		time.Sleep(1 * time.Second)
		duration--
	}
	fmt.Println("\nTimer finished!")
}

func main() {
	var duration int
	fmt.Print("Enter timer duration in seconds: ")
	fmt.Scanf("%d", &duration)
	startTimer(duration)
}
