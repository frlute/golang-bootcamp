package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// question description: https://tour.golang.org/concurrency/7
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walker(t1, ch1)
	go Walker(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	// var ch = make(chan int)
	// go Walker(tree.New(1), ch)
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	//比较两个tree的value是否相等
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
