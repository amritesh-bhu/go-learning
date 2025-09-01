package main

import (
	"fmt"
	"time"
)

func w(from string) {
	for i := range 3 {
		fmt.Println(from, ": ", i)
	}
}

func main() {

	w("direct")

	go w("go routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
