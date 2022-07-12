package main

// https://go.dev/tour/flowcontrol/8

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	k := 0.00001
	for {
		if v := z * z; v < x {
			z = z + k
		} else {
			fmt.Println(v)
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
