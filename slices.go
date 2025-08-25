package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len : ", len(s), "cap :", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	l = s[:5]
	fmt.Println("sl4", l)

	l = s[0:0]
	fmt.Println("sl5", l)

	t1 := []string{"w", "x", "y"}
	fmt.Println("dcl : ", t1)

	t2 := []string{"w", "x", "y"}
	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2")
	}

	twoD := make([][]int, 3)
	for i := range len(twoD) {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := range innerLength {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D ", twoD)
}
