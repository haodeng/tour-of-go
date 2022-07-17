package main

import "fmt"

func do(i interface{}) {
	/**
	A type switch is a construct that permits several type assertions in series.

	A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
	and those values are compared against the type of the value held by the given interface value.
		**/
	switch v := i.(type) {
	// In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i.
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2) // v is 21
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v)) // v is "hello"
		// In the default case (where there is no match), the variable v is of the same interface type and value as i.
	default:
		fmt.Printf("I don't know about type %T!\n", v) // v is bool
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
