package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received messages ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "Hi there"
	select {
	case messages <- msg:
		fmt.Println("Message sent ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message : ", msg)
	case sig := <-signals:
		fmt.Println("Signal received: ", sig)
	default:
		fmt.Println("no activity")
	}
}
