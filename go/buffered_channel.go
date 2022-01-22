package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Limits the channel buff size to 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}