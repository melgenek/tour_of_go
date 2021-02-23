package main

import (
	"./tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	result := false
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			fmt.Println("The values that didn't match", v1, v2, ok1, ok2)
			result = false
			break
		}
		if !ok1 || !ok2 {
			result = true
			break
		}
	}

	return result
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)

	fmt.Println(Same(t1, t2))

}
