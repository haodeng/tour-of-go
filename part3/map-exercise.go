package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		m[v] = m[v] + 1
	}

	fmt.Println(m)
	return m
}

func main() {
	wordCount(" tet test test")
}
