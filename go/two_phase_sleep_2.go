package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Thread struct {
	id       int
	nextTime int
}

func getNextId(id int, n int) int {
	return (id + 1) % n
}

func getLastId(id int, n int) int {
	if id == 0 {
		return n - 1
	} else {
		return id - 1
	}
}

func execute(id int, n int, threads *[5]Thread, mutex sync.Mutex, ch1 chan int, ch2 chan int) {

	// Sleeping
	time_seconds := rand.Intn(10)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - First Time\n", id, time_seconds)
	time.Sleep(sleepTime)

	// Choosing next time
	time_seconds = rand.Intn(10)
	threads[id].nextTime = time_seconds
	fmt.Printf("%d - Next Time = %d\n", id, time_seconds)

	// Waiting other threads
	mutex.Lock()
	count = count - 1
	if count == 0 {
		close(ch1)
		count = n
	}
	mutex.Unlock()
	for range ch1 {}

	// Sleeping Again
	lastId := getLastId(id, n)
	time_seconds = threads[lastId].nextTime
	sleepTime = time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - Second Time\n", id, time_seconds)
	time.Sleep(sleepTime)
	fmt.Printf("%d - Finish\n", id)
	
	// Waiting other threads
	mutex.Lock()
	count = count - 1
	if count == 0 {
		close(ch2)
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
	ch2 := make(chan int)

	for i := 0; i < n; i++ {
		threads[i] = Thread{i, -1}
		go execute(i, n, &threads, mutex, ch1, ch2)
	}
  
	for range ch2 {}
  
	fmt.Println("Finish All")
}
