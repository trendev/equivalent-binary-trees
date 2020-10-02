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

func walk(t *tree.Tree, c chan int) {
	Walk(t, c)
	close(c)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)

	go walk(t1, ch1)
	go walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}

}

func main() {

	t1 := tree.New(1)
	fmt.Println(t1)

	t2 := tree.New(1)
	t2b := &tree.Tree{t2, 12, nil}
	fmt.Println(t2b)

	fmt.Println(Same(t1, t2b))
}
