package main

import (
	"fmt"
	"math"
)

// Sqrt calculates a square root of a float64 via Newton's method.
func Sqrt(x float64) float64 {
	current := x
	actual := math.Sqrt(x)
	near := false
	var i int
	for i = 1; !near; i++ {
		secondTerm := ((current * current) - x) / (2 * current)
		current = current - secondTerm
		if math.Abs(current-actual) < 0.0000000001 {
			near = true
		}
	}
	fmt.Println("Loops taken: ", i-1)
	return current
}

func main() {
	fmt.Println(Sqrt(1234567))
	fmt.Println(math.Sqrt(1234567))
}
