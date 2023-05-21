package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Start the server and listen on a specific port
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port", port)

	client_id := 1

	// Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Printf("Client connected with id %v!\n", client_id)

		// Handle the connection in a separate goroutine
		go handleClient(conn, client_id)

		client_id++
	}
}

func handleClient(conn net.Conn, client_id int) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		// Read the incoming message
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Log the received message
		fmt.Printf("Received message: %q\n", message)

		// Check for an exit condition
		if message == "exit\n" {
			fmt.Printf("Client %v disconnected.\n\nListening for new connections...\n\n", client_id)
			break
		}
	}
}
