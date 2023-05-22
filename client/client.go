package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/JSOD11/Go-Messenger/utils"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	clientId, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return
	}

	fmt.Printf("\nConnected to server with ID %v!\n\n", clientId)

	reader = bufio.NewReader(os.Stdin)

	for {
		// Read input from the user
		fmt.Println("———————————————————————————————————————————————")
		fmt.Printf("Welcome to Messenger! What would you like to do?\n\n")
		fmt.Printf("1 > Login\n2 > Create Account\n3 > List Accounts\n4 > Quit Messenger\n\n")
		fmt.Println("———————————————————————————————————————————————")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		valid, op := utils.ValidateOp(message)
		if !valid {
			fmt.Printf("\nPlease enter 1, 2, 3, or 4.\n\n")
			continue
		}

		// send operation to server
		conn.Write([]byte{op})

		if op == 1 {
			Login()
		} else if op == 2 {
			CreateAccount()
		} else if op == 3 {
			ListAccounts()
		} else if op == 4 {
			fmt.Printf("\nDisconnecting client %v from Messenger...\n\n", clientId)
			break
		}

		// fmt.Printf("\nMessage: %q, %v\n\n", message, len(message))

		// Send the message to the server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println()
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
