package main

import (
	"fmt"
)

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Sequence) modify(i, value int) {
	s[i] = value
}

func (s *Sequence) Len2() int {
	return len(*s)
}

type Stringer interface {
	String2() string
}
type MyFloat float64

func (f MyFloat) String2() string {
	return fmt.Sprintf("MyFloat; %v", f)
}

func main() {
	s := Sequence([]int{2, 3, 5, 7, 11, 13})
	fmt.Println(s.Len())
	fmt.Println(s.Len2())

	fmt.Println("Before swap:", s)
	s.Swap(1, 2)
	fmt.Println("After swap:", s)

	fmt.Println("Before modify:", s)
	s.modify(0, 999)
	fmt.Println("After modify:", s)

	var value interface{} // Value provided by caller.
	//value = "hi"
	value = MyFloat(1.23)
	//value = 1
	switch str := value.(type) {
	case string:
		fmt.Println("string", str)
	case Stringer:
		fmt.Println("Stringer", str.String2())
	default:
		fmt.Println(str)
	}
}
