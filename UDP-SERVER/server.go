package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming UDP packets
	address := ":8080"
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		fmt.Printf("Error starting UDP server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on %s\n", address)

	buffer := make([]byte, 1024)

	for {
		// Read incoming packet
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP connection: %v\n", err)
			continue
		}

		// Print received message
		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", addr, message)

		// Echo the message back to the sender
		_, err = conn.WriteTo([]byte(message), addr)
		if err != nil {
			fmt.Printf("Error writing to UDP connection: %v\n", err)
		}
	}
}
