package main

import "fmt"

func intSeq() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
