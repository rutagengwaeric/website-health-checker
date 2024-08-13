package main

import (
	"fmt"
	"net"
	"time"
)

// Check function to determine if the destination (domain) is reachable on the specified port
func Check(destination string, port string) string {
	// Combine the destination and port to form an address
	address := destination + ":" + port

	// Set a timeout duration of 5 seconds for the connection attempt
	timeout := time.Duration(5 * time.Second)

	// Try to establish a TCP connection to the address with the specified timeout
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string
	if err != nil {
		// If an error occurs, the destination is unreachable; format the status message
		status = fmt.Sprintf("[DOWN] %v is unreachable, \nError: %v", destination, err)
	} else {
		// If no error occurs, the destination is reachable; format the status message
		status = fmt.Sprintf("[UP] %v is reachable, \nFrom: %v\nTo: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	// Return the status message
	return status
}
