package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
	wg.Done()
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
	var wg sync.WaitGroup
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("somekey", &wg)
	}

	wg.Wait() // Wail All Goroutines
	fmt.Println(c.Value("somekey"))
}

/*
1000
*/
