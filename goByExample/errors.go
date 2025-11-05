package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can'[t look with 42]")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("No more tea available!!")
var ErrPower = fmt.Errorf("can't boil water!!")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("Making tea: %w", ErrPower)
	}

	return nil

}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("F failed: ", e)
		} else {
			fmt.Println("F worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should by new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark!!")
			} else {
				fmt.Println("Unknown error: %s\n", err)
			}
		}

		fmt.Println("TEa is ready!")
	}

}
