package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
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

// func main() {

// user1 := User{
// 	Name: "Yudha",
// }

// user2 := User{
// 	Name: "Sugiharto",
// }

// room := Room{
// 	Participant: &[]User{user1, user2},
// 	Chat:        make(chan Message, 10),
// }

// }

func WebSocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			return
		}

		fmt.Println("Received message:", message)

		err = websocket.Message.Send(ws, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(WebSocketHandler))
	fmt.Println("WebSocket server listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
