package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func execute(id int, wg *sync.WaitGroup) {

	// Sleeping
	time_seconds := rand.Intn(10)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - First Time\n", id, time_seconds)
	time.Sleep(sleepTime)
	fmt.Printf("%d - Finish\n", id)
	wg.Done()
}

const n = 5
var count = n

func main() {

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go execute(i, &wg)
	}
	
	wg.Wait()
	
	fmt.Println("Finish All")
	
}

