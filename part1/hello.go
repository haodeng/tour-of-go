package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi) // a name is exported if it begins with a capital letter

}
