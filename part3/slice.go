package main

import (
	"fmt"
	"strings"
)

func slicePointer() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// A slice literal is like an array literal without the length.
func sliceliterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// you may omit the high or low bounds to use their defaults instead.
//The default is zero for the low bound and the length of the slice for the high bound.
/**
For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]
**/
func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

/**
The length of a slice is the number of elements it contains.
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
**/
func sliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//The zero value of a slice is nil.
//A nil slice has a length and capacity of 0 and has no underlying array.
func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {
	// The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSlice2("a", a) // a len=5 cap=5 [0 0 0 0 0]

	// To specify a capacity, pass a third argument
	b := make([]int, 0, 5)
	printSlice2("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Slices can contain any type, including other slices.
func sliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/**
func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
**/
func appendToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// An array has a fixed size. A slice, on the other hand, is a dynamically-sized
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// The type []T is a slice with elements of type T.
	// creates a slice which includes elements 1 through 3
	var s []int = primes[1:4] // [3 5 7]
	fmt.Println(s)

	slicePointer()

	sliceliterals()

	sliceDefaults()

	sliceLengthCapacity()

	nilSlice()

	makeSlice()

	sliceOfSlice()

	appendToSlice()
}
