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

	var clientId byte = 1

	// Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Printf("Client connected with id %v!\n", clientId)

		// Handle the connection in a separate goroutine
		go handleClient(conn, clientId)

		clientId++
	}
}

func handleClient(conn net.Conn, clientId byte) {

	defer conn.Close()

	// send client its own ID
	conn.Write([]byte{clientId})

	reader := bufio.NewReader(conn)

	for {
		// Read the incoming operation client has chosen
		op, err := reader.ReadByte()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// receive operation from client
		if op == 1 {
			Login()
		} else if op == 2 {
			CreateAccount()
		} else if op == 3 {
			ListAccounts()
		} else if op == 4 {
			fmt.Printf("Client %v disconnected.\n\nListening for new connections...\n\n", clientId)
			return
		}
	}
}

func Login() byte {
	return 1
}

func CreateAccount() byte {
	return 1
}

func ListAccounts() byte {
	return 1
}
