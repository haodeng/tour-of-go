package main

import (
	"fmt"
)

func pic(dx, dy int) [][]uint8 {
	var s [][]uint8
	for i := 0; i < dy; i++ {
		var s2 []uint8
		for j := 0; j < dx; j++ {
			s2 = append(s2, uint8(i+j))
		}
		s = append(s, s2)
	}

	return s
}

func main() {
	s := pic(2, 3)
	fmt.Println(s)
}
