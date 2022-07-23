package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (lst *List[T]) add(v T) {

	for lst.next != nil {
		lst = lst.next
	}

	lst.next = &List[T]{val: v}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	t := List[int]{val: 1}
	t.next = &List[int]{val: 2}

	fmt.Println(t.val)
	fmt.Println(t.next.val)

	lst := List[int]{val: 1}
	lst.add(10)
	lst.add(13)
	lst.add(23)
	fmt.Println("list:", lst.GetAll())
}
