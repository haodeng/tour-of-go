package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Unlike other languages like Java, you don’t need to explicitly specify that a type implements an interface using something like an implements keyword.
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var i2 = T{"hello2"}
	i2.M()

	T{"hello3"}.M()
}
