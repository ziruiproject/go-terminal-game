package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
}

type Room struct {
	Participant *[]User
	Chat        chan Message
}

type Message struct {
	From    *User
	Message string
}

func (user *User) sendMessage(room Room, msg string) {
	room.Chat <- Message{
		From:    user,
		Message: msg,
	}
}

func main() {

	user1 := User{
		Name: "Yudha",
	}

	user2 := User{
		Name: "Sugiharto",
	}

	room := Room{
		Participant: &[]User{user1, user2},
		Chat:        make(chan Message, 10),
	}

	fmt.Println(room)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read and process data from the client
	// ...

	// Write data back to the client
	// ...
}
