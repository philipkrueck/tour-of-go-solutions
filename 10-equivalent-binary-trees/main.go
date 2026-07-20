// Video explanation: https://www.youtube.com/watch?v=Yl6ix6lhtFs&t=8s

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan<- int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan<- int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, open1 := <-ch1
		v2, open2 := <-ch2

		if !open1 || !open2 {
			return open1 == open2
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	t1, t2 := tree.New(1), tree.New(1)

	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("==:", Same(t1, t2))
}
