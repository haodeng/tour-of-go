package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/**
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.
**/
//  the Abs method has a receiver of type Vertex named v.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Remember: a method is just a function with a receiver argument.
//Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// You can declare a method on non-struct types, too.
//a numeric type MyFloat with an Abs method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	// call method on Vertex struct
	fmt.Println(v.Abs())

	// call regular function
	fmt.Println(Abs(v))

	// call method on MyFloat
	myFloat := MyFloat(-math.Sqrt2)
	fmt.Println(myFloat.Abs())
}
