package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

// TestClockServer tests if the clock server sends time data to a client.
func TestClockServer(t *testing.T) {
	// Start the server in a goroutine
	go func() {
		main()
	}()

	// Allow the server time to start
	time.Sleep(1 * time.Second)

	// Connect to the server as a client
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Read the time data from the server
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read from server: %v", err)
	}

	// Check if the received string is in the correct time format
	line = strings.TrimSpace(line)
	_, err = time.Parse("15:04:05", line)
	if err != nil {
		t.Errorf("Received invalid time format: %s", line)
	}
}
