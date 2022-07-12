package main

import "fmt"

// The var statement declares a list of variables; as in function argument lists, the type is last.
// package level var
var c, python, java bool

// Variables with initializers
var i, j int = 1, 2

func main() {
	var i2 int
	fmt.Println(i2, c, python, java)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var c2, python2, java2 = true, false, "no!"
	fmt.Println(i, j, c2, python2, java2)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	k := 3
	fmt.Println(k)
}
