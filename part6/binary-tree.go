package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	if t.Right != nil {
		Walk(t.Right, ch)
	}
	ch <- t.Value

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	s1 := treeToSortedSlice(t1)
	s2 := treeToSortedSlice(t2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < 10; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func treeToSortedSlice(t *Tree) []int {
	c := make(chan int, 10)
	go Walk(t, c)

	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, <-c)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return s
}

func main() {
	c := make(chan int, 10)
	go Walk(New(1), c)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Same: ", Same(New(1), New(1)))
	fmt.Println("Same: ", Same(New(1), New(2)))
}
