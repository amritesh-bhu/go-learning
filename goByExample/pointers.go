package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0
}

func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 2
	fmt.Println("initial: ", i)

	zeroVal(i)
	fmt.Println("zeroVal: ", i)

	zeroPtr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)

}
