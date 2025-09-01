package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item
}

func main() {
	instack := Stack[int]{}
	instack.Push(10)
	instack.Push(20)
	instack.Push(30)

	fmt.Println("Push", instack)

	instack.Pop()
	fmt.Println("Pop: ", instack)
}
