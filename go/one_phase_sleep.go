package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Thread struct {
	id       int
}

func execute(id int, n int, threads *[5]Thread, mutex sync.Mutex, ch1 chan int) {

	// Sleeping
	time_seconds := rand.Intn(10)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - First Time\n", id, time_seconds)
	time.Sleep(sleepTime)
	fmt.Printf("%d - Finish\n", id)
	
	// Waiting other threads
	mutex.Lock()
	count = count - 1
	if count == 0 {
		close(ch1)
		count = n
	}
	mutex.Unlock()
}

const n = 5
var count = n

func main() {
	
	var threads [n]Thread
	var mutex sync.Mutex
	ch1 := make(chan int)

	for i := 0; i < n; i++ {
		threads[i] = Thread{i}
		go execute(i, n, &threads, mutex, ch1)
	}
	
	for range ch1 {}
	
	fmt.Println("Finish All")
	
}
