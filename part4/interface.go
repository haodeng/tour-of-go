package main

import (
	"fmt"
	"math"
)

/**
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.
**/
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
	// (value, type)
	describe(a) // (-1.4142135623730951, main.MyFloat)

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	describe(a) // (&{3 4}, *main.Vertex)

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v  //Compile error
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
