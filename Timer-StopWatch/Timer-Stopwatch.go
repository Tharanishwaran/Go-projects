package main

import (
	"fmt"
	"time"
)



func main() {

   var option string

   fmt.Println("Enter A Option for Timer | Enter B option for Stopwatch")

   fmt.Scanln(&option)

   if option == "A" {

	var duration int
	fmt.Print("Enter timer duration in seconds: ")
	fmt.Scanf("%d", &duration)
	startTimer(duration)
	
  }else if option == "B" {

   var command string
	fmt.Print("Type 'start' to begin the stopwatch: ")
	fmt.Scanf("%s", &command)
	if command == "start" {
		go startStopwatch()
		fmt.Println("\nStopwatch started. Type 'stop' to end.")
		for {
			fmt.Print("\nType 'stop' to end: ")
			fmt.Scanf("%s", &command)
			if command == "stop" {
				break
			}
		}
	}
	fmt.Println("Stopwatch stopped.")
}else{

  fmt.Println("you enter invalid options")

}

}




func startTimer(duration int) {
	for duration > 0 {
		fmt.Printf("\rTime remaining: %d seconds", duration)
		time.Sleep(1 * time.Second)
		duration--
	}
	fmt.Println("\nTimer finished!")
}

func startStopwatch() {
	startTime := time.Now()
	for {
		select {
		case <-time.After(1 * time.Second):
			elapsed := time.Since(startTime)
			fmt.Printf("\rElapsed time: %v", elapsed)
		}
	}
}