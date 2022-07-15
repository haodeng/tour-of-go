package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//v Vertex, value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
// Methods with pointer receivers can modify the value to which the receiver points
//v *Vertex, pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// regular functions
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// method
	v := Vertex{3, 4}
	v.Scale(10) // Go interprets the statement v.Scale(10) as (&v).Scale(10) since the Scale method has a pointer receiver.
	fmt.Println(v.Abs(), v.X, v.Y)

	// method
	p := &Vertex{4, 3}
	p.Scale(10)
	fmt.Println(p.Abs(), p.X, p.Y)
	// function
	Scale(p, 10)
	fmt.Println(p.Abs(), p.X, p.Y)

	// function
	v2 := Vertex{3, 4}
	Scale(&v2, 10)
	fmt.Println(Abs(v2))
}
