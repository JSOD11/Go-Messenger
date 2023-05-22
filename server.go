package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/JSOD11/Go-Messenger/utils"
)

type UserManager struct {
	idCounter byte
	users     map[string]*User
}

type User struct {
	username       string
	isLoggedIn     bool
	unreadMessages []string
}

func main() {

	um := UserManager{
		idCounter: 1,
		users:     make(map[string]*User), // users[username] -> User struct
	}

	// Start the server and listen on a specific port
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Printf("Server is listening on port %v\n", port)

	// Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Printf("Client connected with id %v!\n", um.idCounter)

		// Handle user connection in a separate goroutine
		go um.handleClient(conn, um.idCounter)

		um.idCounter++
	}
}

func (um *UserManager) handleClient(conn net.Conn, clientId byte) {

	defer conn.Close()

	// send client its own ID
	conn.Write([]byte{clientId})

	connReader := bufio.NewReader(conn)

	for {
		// Read the incoming operation client has chosen
		op, err := connReader.ReadByte()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		if op == 1 {
			um.login(conn, connReader)
		} else if op == 2 {
			um.createAccount(conn, connReader)
		} else if op == 3 {
			um.listAccounts(conn)
		} else if op == 4 {
			fmt.Printf("Client %v disconnected.\n\nListening for new connections...\n\n", clientId)
			return
		}
	}
}

func (um *UserManager) login(conn net.Conn, connReader *bufio.Reader) {
	username, err := connReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	username = username[0 : len(username)-1]
	fmt.Printf("\nNew login attempt: %v : ", username)

	if user, ok := um.users[username]; ok && !user.isLoggedIn {
		// login attempt is valid
		fmt.Print("SUCCESS\n\n")
		user.isLoggedIn = true
		conn.Write([]byte{utils.SUCCESS})
		um.logAccounts()
		um.userMenu(conn, connReader, username)
	} else {
		// name does not exist in server, invalid
		fmt.Print("FAILURE\n\n")
		conn.Write([]byte{utils.FAILURE})
	}
}

func (um *UserManager) createAccount(conn net.Conn, connReader *bufio.Reader) {
	username, err := connReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	username = username[0 : len(username)-1]

	if _, ok := um.users[username]; ok {
		// name already taken
		conn.Write([]byte{utils.FAILURE})
	} else {
		// created account successfully
		um.users[username] = new(User)
		um.users[username].username = username
		conn.Write([]byte{utils.SUCCESS})
	}

	um.logAccounts()
}

func (um *UserManager) listAccounts(conn net.Conn) {
	var accountNames []byte
	for _, user := range um.users {
		// client will know usernames are split by \n
		accountNames = append(accountNames, []byte(user.username)...)
		accountNames = append(accountNames, '\n')
	}
	// client will scan for $ which indicates end of message
	accountNames = append(accountNames, '$')
	fmt.Printf("\nlistAccounts message: %q\n\n", accountNames)
	conn.Write(accountNames)
}

func (um *UserManager) logAccounts() {
	fmt.Println("——————————————")
	if len(um.users) == 0 {
		fmt.Printf("No accounts on server")
	} else {
		if len(um.users) == 1 {
			fmt.Printf("1 account: \n")
		} else {
			fmt.Printf("%v account(s): \n", len(um.users))
		}
		for _, user := range um.users {
			fmt.Printf("%v | %v | %v\n", user.username, user.isLoggedIn, user.unreadMessages)
		}
	}
	fmt.Println()
}

func (um *UserManager) userMenu(conn net.Conn, connReader *bufio.Reader, username string) {
	for {
		// Read the incoming operation client has chosen
		op, err := connReader.ReadByte()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		if op == 1 {
			um.routeMessage()
		} else if op == 2 {
			um.showMessages()
		} else if op == 3 {
			um.users[username].isLoggedIn = false
			fmt.Printf("%v logged out\n", username)
			um.logAccounts()
			break
		} else if op == 4 {
			delete(um.users, username)
			fmt.Printf("Deleted account: %v\n", username)
			um.logAccounts()
			break
		}
	}
}

func (um *UserManager) routeMessage() {
	fmt.Println("Route message")
}

func (um *UserManager) showMessages() {
	fmt.Println("Show messages")
}
