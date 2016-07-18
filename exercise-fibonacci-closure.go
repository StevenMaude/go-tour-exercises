package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n, current, previous := 0, 0, 0

	return func() int {
		v := current
		if n == 0 || n == 1 {
			v = n
		} else {
			v = current + previous
		}
		previous, current = current, v
		n++
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
