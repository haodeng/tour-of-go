package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

/**
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception,
but in Go it is common to write methods that gracefully handle being called with a nil receiver
**/
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t // nil value
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	// nil interface
	/**
	A nil interface value holds neither value nor concrete type.
	Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	**/
	var i2 I
	describe(i2) // (<nil>, <nil>)
	//i2.M()  // runtime error: invalid memory address or nil pointer dereference

	emptyInterface()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
