package main

import "fmt"

// A struct is a collection of fields.
type Vertex struct {
	X    int
	Y    int
	Desc string
}

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
var (
	v1 = Vertex{1, 2, "test"}   // has type Vertex
	v2 = Vertex{X: 1}           // Y:0 is implicit, You can list just a subset of fields by using the Name: syntax
	v3 = Vertex{}               // X:0 and Y:0
	p2 = &Vertex{1, 2, "test2"} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2, "test"})

	v := Vertex{1, 2, "test2"}
	v.X = 4
	// Struct fields are accessed using a dot.
	fmt.Println(v.X)
	fmt.Println(v.Y)

	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p := &v
	(*p).X = 1e9
	fmt.Println(v)
	p.X = 1
	fmt.Println(v)

	fmt.Println("Struct Literals test:", v1, v2, v3, p2)
}
