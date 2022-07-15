package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//The range form of the for loop iterates over a slice or map.
func main() {
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	// You can skip the index or value by assigning to _.
	for i, _ := range pow {
		fmt.Printf("index %d\n", i)
	}

	// If you only want the index, you can omit the second variable.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("value %d\n", value)
	}
}
