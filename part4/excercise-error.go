package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

/**
A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
**/
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	k := 0.00001
	for {
		if v := z * z; v < x {
			z = z + k
		} else {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	i, err := Sqrt(-2)
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(i)
	}
}
