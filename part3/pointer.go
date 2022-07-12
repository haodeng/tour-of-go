package main

import "fmt"

func main() {
	i, j := 42, 2701

	// A pointer holds the memory address of a value.
	// The & operator generates a pointer to its operand.
	p := &i // point to i
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // read i through the pointer: 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // i: 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // 73

	// The type *T is a pointer to a T value. Its zero value is nil.
	var p2 *int
	i2 := 42
	p2 = &i2
	fmt.Println(*p2) // 42
}
