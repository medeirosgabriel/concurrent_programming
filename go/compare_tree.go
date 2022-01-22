package main

import "fmt"

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func Walk(t *Tree, n int, ch chan int) {
	
	if ((*t.Left) != Tree{}) {
		Walk(t.Left, n + 1, ch)
	}
	
	ch <- t.Value
	
	if ((*t.Right) != Tree{}) {
		Walk(t.Right, n + 1, ch)
	}
	
	if (n == 0) {
		close(ch)	
	}
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, 0, ch1)
    go Walk(t2, 0, ch2)
	
	for value1 := range ch1 {
		value2 := <-ch2
		fmt.Printf("%d - %d\n", value1, value2)
		if (value1 != value2) {
			return false	
		}
	}
	
	return true
}

func main() {
	t1 := &Tree{&Tree{&Tree{}, 2, &Tree{}}, 4, &Tree{&Tree{}, 5, &Tree{}}}
	t2 := &Tree{&Tree{&Tree{&Tree{}, 2, &Tree{}}, 4, &Tree{}}, 5, &Tree{}}
	fmt.Println(Same(t1, t2))
}
