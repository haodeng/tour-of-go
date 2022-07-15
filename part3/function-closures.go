package main

import "fmt"

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

// the adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pre1 := 0
	pre2 := 0
	fib := 0
	i := 0
	return func() int {
		if i == 0 {
			fib = 0
		} else if i == 1 {
			fib = 1
		} else {
			pre2 = pre1
			pre1 = fib
			fib = pre1 + pre2
		}
		i += 1

		return fib
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v,", f())
	}
}
