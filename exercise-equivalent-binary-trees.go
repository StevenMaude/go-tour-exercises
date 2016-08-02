// The official solution here:
// https://github.com/golang/tour/blob/c9941e54e5b8e9618a8c951bc89798f85f6a7a71/solutions/binarytrees.go
// is much better than this one.

// 1. It uses a wrapper function that just calls a function similar to
// Walk() here, and then closes the channel.
// It means you don't need to hardcode to 10 values when printing, since you
// can iterate over range of the channel to print all the tree values.

// (Closing inside Walk() doesn't work correctly as one call will reach
// the end before it's finished.)

// Likewise for checking value equality: return False if values from the
// channels are unequal while iterating over them in an infinite for loop.
// Stop when you run out of values (an ok value is false) and then compare
// the ok values; if they're both false when they're the same sequence. If
// one's still open, then they can't be the same sequence.

// 2. The Walk() there also only checks the current t == nil; you
// don't actually need to check the branches individually, you
// can follow all the way into a nil leaf, then just return if you hit
// one. Walk() on that branch then ends immediately, we return back up, and
// pass the t.Value for the valid parent into the channel.

// 3. They make a further improvement here:
// https://github.com/golang/tour/blob/c9941e54e5b8e9618a8c951bc89798f85f6a7a71/solutions/binarytrees_quit.go
// where they use another channel of ints to signal to Walk() to just
// return once the comparison is finished, so Walk() doesn't continue to run
// unnecessarily, when two trees are of different lengths.
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 1; i <= 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	// Should be 1, 2... 10.
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

	t1 := tree.New(1)
	t2 := tree.New(1)
	// Should be true.
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(t1, t2))

	t1 = tree.New(1)
	t2 = tree.New(2)
	// Should be false.
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(t1, t2))
}
