package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	t := List[int]{val: 1}
	t.next = &List[int]{val: 2}

	fmt.Println(t.val)
	fmt.Println(t.next.val)
}
