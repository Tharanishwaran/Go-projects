package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Resolve server address
	address := "localhost:8080"
	serverAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Printf("Error resolving server address: %v\n", err)
		return
	}

	// Dial UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Error dialing UDP connection: %v\n", err)
		return
	}
	defer conn.Close()

	// Send message to server
	message := "Hello, UDP server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to UDP connection: %v\n", err)
		return
	}
	fmt.Printf("Sent message to server: %s\n", message)

	// Receive response from server
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Printf("Error reading from UDP connection: %v\n", err)
		return
	}

	// Print received response
	response := string(buffer[:n])
	fmt.Printf("Received response from %s: %s\n", addr, response)
}
