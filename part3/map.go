package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// Map literals are like struct literals, but the keys are required.
func mapLiteral() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
}

/**
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.
**/
func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	v = m["Answer"]
	fmt.Println("The value:", v)

	// use the blank identifier (_) in place of the usual variable for the value.
	_, ok = m["Answer"]
	fmt.Println("Present?", ok)
}

func main() {
	//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	//The make function returns a map of the given type, initialized and ready for use.
	fmt.Println(m) // nil map
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	mapLiteral()
	mutatingMaps()
}
