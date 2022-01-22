package main

import "fmt"

/*
The select statement allows a wait in the goroutine on multiple communication operations.
The select block waits until one of its cases can execute, then it executes that case. He chooses one at random if several are ready.
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
0
1
1
2
3
5
8
13
21
34
quit
*/