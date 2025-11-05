package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectngl struct {
	height, width float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

func (r rectngl) area() float64 {
	return r.height * r.width
}

func (r rectngl) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perim() float64 {
	hyp := math.Sqrt(t.base*t.base + t.height*t.height)
	return hyp + t.base + t.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimeter: ", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle radius: ", c.radius)
	}
}

func main() {
	r := rectngl{height: 5, width: 5}
	c := circle{radius: 5}
	t := triangle{base: 3, height: 4}

	measure(r)
	measure(c)
	measure(t)

	detectCircle(r)
	detectCircle(c)
	detectCircle(t)
}
