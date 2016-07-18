package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt calculates a square root of a float64 via Newton's method.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	current := x
	var i int
	for i = 1; i <= 10; i++ {
		secondTerm := ((current * current) - x) / (2 * current)
		current = current - secondTerm
	}
	fmt.Println("Loops taken: ", i-1)
	return current, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
