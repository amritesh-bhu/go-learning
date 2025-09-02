package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 3 {
		select {
		case msg1 := <-c1:
			fmt.Println("Received: ", msg1)
		case msg1 := <-c2:
			fmt.Println("Received: ", msg1)
		}
	}
}
