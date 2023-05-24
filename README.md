# Go-Messenger

Rebuild of a distributed messenger system I built in CS 262 for fun and as I learn Go. See JSOD11/CS262-Messenger for the Python build.

<img width="1220" alt="Screenshot 2023-05-24 at 3 26 02 PM" src="https://github.com/JSOD11/Go-Messenger/assets/55005116/a09b5080-5e59-45c8-ab45-e98c2a538184">

# Usage

## Server
In the `go-messenger` directory, start the server with
```
go run .
```
This begins a server that will log actions taken by users.

## Client
In a separate terminal, switch into the `go-messenger/client` directory and start a client with
```
go run .
```
This will pull up a menu with options to login (1), create an account (2), list accounts (3), or quit (4). To take one of these actions, enter the corresponding number and hit enter in the terminal. After creating an account and logging in, a user can send messages to other users (1), view their own messages (2), log out (3), or delete their account (4). Multiple clients can connect to the server concurrently. The system is designed such that duplicate accounts cannot be created. While one client is logged in to a user account, other clients will not be able to log in to that account. If a user deletes their account, all their unread messages are lost. The data is stored within the server and does not persist when the server is shut down.
