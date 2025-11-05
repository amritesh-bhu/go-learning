package main

import "fmt"

func varidaicSum(args ...int) {
	fmt.Println(args, " ")
	total := 0

	for _, num := range args {
		total += num
	}

	fmt.Println(total)
}

func main() {
	varidaicSum(1, 2)
	varidaicSum(1, 2, 3)
	s := []int{4, 5, 6, 7}
	varidaicSum(s...)
}
