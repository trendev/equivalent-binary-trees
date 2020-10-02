package main

import (
	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch and closes the channel
func Walk(t *tree.Tree, c chan int) {
	walk(t, c)
	close(c)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// if channels cap() are different
		// ok1 or ok2 will be true (and the opposite)
		// and v1 or v2 will be the default value 0
		if !ok1 && !ok2 { // end of channel
			return true
		}

		if v1 != v2 {
			return false
		}
	}

}

func main() {
	// DO NOTHING
}
