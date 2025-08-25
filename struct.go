package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "fred"})
	fmt.Println(newPerson("jon"))
	fmt.Println(&person{"amrit", 30})
}
