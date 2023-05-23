package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/JSOD11/Go-Messenger/utils"
)

func main() {
	// connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	// receive clientId from server
	connReader := bufio.NewReader(conn) // connReader reads from server
	clientId, err := connReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return
	}
	fmt.Printf("\nConnected to server with ID %v!\n\n", clientId)

	inputReader := bufio.NewReader(os.Stdin) // inputReader reads from terminal
	for {                                    // main loop
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
			login(conn, connReader, inputReader)
		} else if op == 2 {
			createAccount(conn, connReader, inputReader)
		} else if op == 3 {
			listAccounts(connReader, inputReader)
		} else if op == 4 {
			fmt.Printf("\nDisconnecting client %v from Messenger...\n\n", clientId)
			break
		}
	}
}

func login(conn net.Conn, connReader *bufio.Reader, inputReader *bufio.Reader) {
	fmt.Println("Messenger Login")
	fmt.Println("———————————————————————————————————————————————")
	fmt.Printf("Please enter a username to log in: \n\n")

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
		//fmt.Printf("Successfully logged in with username %v!\n", username[0:len(username)-1])
		userMenu(conn, connReader, inputReader, username[0:len(username)-1])
	} else {
		fmt.Printf("User does not exist or is alread logged in. Please try again or create a new account.\n")
	}
}

func userMenu(conn net.Conn, connReader *bufio.Reader, inputReader *bufio.Reader, username string) {
	for {
		// Read input from the user
		fmt.Println("———————————————————————————————————————————————")
		fmt.Printf("Account: %v\n\n", username)
		fmt.Printf("1 > Send messages\n2 > View my messages\n3 > Logout\n4 > Delete account\n\n")
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
			sendMessage(conn, connReader, inputReader)
		} else if op == 2 {
			viewMessages(connReader, inputReader)
		} else if op == 3 { // log out
			fmt.Printf("%v logged out\n\n", username)
			break
		} else if op == 4 { // delete account
			fmt.Printf("%v's account deleted successfully\n\n", username)
			break
		}
	}
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

func sendMessage(conn net.Conn, connReader *bufio.Reader, inputReader *bufio.Reader) {
	fmt.Println("———————————————————————————————————————————————")
	fmt.Printf("Who would you like to send messages to?\n\n")
	target, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	utils.ResetScreen()

	// send target username to server
	conn.Write([]byte(target))

	result, err := connReader.ReadByte()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if result == utils.SUCCESS {
		fmt.Println("———————————————————————————————————————————————")
		fmt.Printf("Sending messages to %v. Enter \\E to exit.\n\n", target[0:len(target)-1])
		for {
			message, err := inputReader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			//fmt.Printf("%q\n", message)
			conn.Write([]byte(message))
			if message == "\\E\n" {
				break
			}
		}
		utils.ResetScreen()
	} else {
		fmt.Printf("The username you provided does not exist! Please try again.\n")
	}
}

func viewMessages(connReader *bufio.Reader, inputReader *bufio.Reader) {
	scanner := bufio.NewScanner(connReader)
	scanner.Split(bufio.ScanBytes)

	fmt.Printf("\nUnread messages\n")
	fmt.Printf("———————————————————————————————————————————————\n")

	var message []byte
	var char byte
	// read message byte by byte
	for scanner.Scan() {
		char = scanner.Bytes()[0]
		if char == '$' { // $ indicates end of unread messages
			break
		} else if char == '\n' { // \n indicates end of accountName
			fmt.Println(string(message))
			message = []byte{}
		} else {
			message = append(message, char)
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
