package main

import (
	"fmt"
	"os"
)

func main() {
	//Formatted printing in Go: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on.
	fmt.Printf("Hello %d\n", 23)

	// If you don't need to provide a format string.
	// For each of Printf, Fprintf and Sprintf there is another pair of functions, for instance Print and Println.
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	printStruct()

	m := MyString("my")
	fmt.Println(m)
}

type T struct {
	a int
	b float64
	c string
}

func printStruct() {
	t := &T{7, -2.35, "abc\tdef"}
	// format %v (for “value”); the result is exactly what Print and Println would produce.
	fmt.Printf("%v\n", t) // &{7 -2.35 abc   def}

	// format %+v annotates the fields of the structure with their names
	fmt.Printf("%+v\n", t) // &{a:7 b:-2.35 c:abc     def}

	// and for any value the alternate format %#v prints the value in full Go syntax.
	fmt.Printf("%#v\n", t) // &main.T{a:7, b:-2.35, c:"abc\tdef"}

	// %T prints the type of a value.
	fmt.Printf("%T\n", t) // *main.T
	fmt.Printf("%T\n", 1) // int
}

type MyString string

func (m MyString) String() string {
	// return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
	return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
