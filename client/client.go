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

	connReader := bufio.NewReader(conn)
	clientId, err := connReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return
	}

	fmt.Printf("\nConnected to server with ID %v!\n\n", clientId)

	inputReader := bufio.NewReader(os.Stdin)

	for {
		// Read input from the user
		fmt.Println("———————————————————————————————————————————————")
		fmt.Printf("Welcome to Messenger! What would you like to do?\n\n")
		fmt.Printf("1 > Login\n2 > Create Account\n3 > List Accounts\n4 > Quit Messenger\n\n")
		fmt.Println("———————————————————————————————————————————————")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		utils.ResetScreen()

		valid, op := utils.ValidateOp(input)
		if !valid {
			fmt.Printf("\nPlease enter 1, 2, 3, or 4.\n\n")
			continue
		}

		// send operation to server
		conn.Write([]byte{op})

		if op == 1 {
			login()
		} else if op == 2 {
			createAccount(conn, connReader, inputReader)
		} else if op == 3 {
			listAccounts(connReader, inputReader)
		} else if op == 4 {
			fmt.Printf("\nDisconnecting client %v from Messenger...\n\n", clientId)
			break
		}

		// Send the message to the server
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println()
	}
}

func login() byte {
	return 1
}

func createAccount(conn net.Conn, connReader *bufio.Reader, inputReader *bufio.Reader) {
	fmt.Println("Create Account")
	fmt.Println("———————————————————————————————————————————————")
	fmt.Printf("Please enter a username: \n\n")

	username, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	utils.ResetScreen()

	// send username to server for processing
	conn.Write([]byte(username))

	// server sends back success or failure
	result, err := connReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if result == utils.SUCCESS {
		fmt.Printf("Successfully created account with username %v!\n", username[0:len(username)-1])
	} else {
		fmt.Printf("The username you provided has already been taken. Please try again.\n")
	}

}

func listAccounts(connReader *bufio.Reader, inputReader *bufio.Reader) {
	scanner := bufio.NewScanner(connReader)
	scanner.Split(bufio.ScanBytes)

	fmt.Printf("\nAccounts registered on server\n")
	fmt.Printf("———————————————————————————————————————————————\n")

	var accountName []byte
	var char byte
	// read message byte by byte
	for scanner.Scan() {
		char = scanner.Bytes()[0]
		if char == '$' { // $ indicates end of message
			break
		} else if char == '\n' { // \n indicates end of accountName
			fmt.Println(string(accountName))
			accountName = []byte{}
		} else {
			accountName = append(accountName, char)
		}
	}

	fmt.Printf("\nPress enter to return to the main menu\n\n")
	_, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	utils.ResetScreen()
}
