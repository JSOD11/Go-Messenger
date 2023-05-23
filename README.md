# Go-Messenger

Rebuild of a distributed messenger system I built in CS 262 for fun and as I learn Go. See JSOD11/CS262-Messenger for the Python build.

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
This will pull up a menu with options to login, create an account, list accounts, or quit. After creating an account and logging in, a user can send messages to other users, view their own messages, log out, or delete their account. Multiple clients can connect to the server concurrently.
