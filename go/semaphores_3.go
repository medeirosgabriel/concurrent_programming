package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
    count int
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, ch chan int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.count = c.count - 1
	if (c.count == 0) {
	    fmt.Println("i am the last thread")
	    close(ch)
	} else {
	    //fmt.Println("i am not the last thread")
	}
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock() // mutex.down()
	// Lock so only one goroutine at a time can access the map c.v.
	key_value := c.v[key]
	defer c.mu.Unlock() // mutex.up()
	return key_value
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	n := 1000
	c.count = n
	ch := make(chan int)
	for i := 0; i < n; i++ {
		go c.Inc("somekey", ch)
	}
	for range ch {}
	fmt.Println(c.Value("somekey"))
}

/*
i am the last thread
1000
*/
