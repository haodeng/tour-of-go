package main

import (
	"fmt"
	"math/cmplx"
)

/**
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
**/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var ui8 uint8 = 1
	fmt.Printf("Type: %T Value: %v\n", ui8, ui8)

	// Zero values
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
