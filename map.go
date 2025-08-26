package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)
	fmt.Println("map ", m)

	m["k1"] = 7
	m["k2"] = 9
	fmt.Println("map with keys: ", m)

	v1 := m["k1"]
	fmt.Println("value : ", v1)

	v2 := m["k3"]
	fmt.Println("value for key not defined : ", v2)

	fmt.Println("length of map : ", len(m))

	delete(m, "k2")
	fmt.Println("after delete : ", m)

	clear(m)
	fmt.Println("after clear : ", m)

	_, prs := m["k2"]
	fmt.Println("prs : ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map : ", n)

	n1 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n1) {
		fmt.Println("n == n1")
	}

}
