package main

import "fmt"

func main() {
	// The type [n]T is an array of n values of type T.
	// cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println("length:", len(a))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// array
	a2 := [...]string{"no error", "Eio", "invalid argument", "test"}
	fmt.Println(a2)
	fmt.Printf("type: %T\n", a2) // [4]string

	// slice, not array
	s := []string{"no error", "Eio", "invalid argument"}
	fmt.Println(s)
	fmt.Printf("type: %T\n", s) // []string

	//map
	m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
	fmt.Printf("type: %T\n", m) // map[int]string
}
