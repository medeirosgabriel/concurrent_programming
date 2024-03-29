﻿package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	s := fmt.Sprintf("Started %d", cap(c))
	fmt.Println(s)
	go fibonacci(cap(c), c)
	for i := range c { // Stop when the channel is closed
		fmt.Println(i)
	}
}

/*
Started 10
0
1
1
2
3
5
8
13
*/