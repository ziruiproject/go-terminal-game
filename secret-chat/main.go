package main

import (
	"fmt"
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

	go func() {
		user1.sendMessage(room, "Halo Gan")
		user2.sendMessage(room, "Halo juga gan")
		close(room.Chat)
	}()

	for data := range room.Chat {
		fmt.Println("Message: " + data.Message)
	}
}
