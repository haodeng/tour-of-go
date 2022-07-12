package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// The init and post statements are optional.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	// For is Go's "while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//  omit the loop condition it loops forever
	//for {
	//	fmt.Println("never ends")
	//}
}
