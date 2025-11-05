package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (ae *argError) Error() string {
	return fmt.Sprintf("%s - %d ", ae.message, ae.arg)
}

func g(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	_, err := g(43)

	var e *argError

	if errors.As(err, &e) {
		fmt.Println(e.arg)
		fmt.Println(e.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
