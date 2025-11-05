package main

import "fmt"

type rect struct {
	height, width int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perm() int {
	return 2*r.height + 2*r.width
}

func main() {

	r := rect{height: 4, width: 5}

	fmt.Println("area : ", r.area())
	fmt.Println("perimeter : ", r.perm())

	rp := &r

	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perm())

}
