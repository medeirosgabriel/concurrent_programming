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

func execute(id int, n int, threads *[5]Thread, wg1 *sync.WaitGroup, wg2 *sync.WaitGroup) {

	// Sleeping
	time_seconds := rand.Intn(10)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - First Time\n", id, time_seconds)
	time.Sleep(sleepTime)

	// Choosing next time
	time_seconds = rand.Intn(10)
	threads[id].nextTime = time_seconds
	fmt.Printf("%d - Choose Time = %d\n", id, time_seconds)

	// Waiting other threads
	wg1.Done()
	wg1.Wait()

	// Sleeping Again
	lastId := getLastId(id, n)
	time_seconds = threads[lastId].nextTime
	sleepTime = time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - Second Time\n", id, time_seconds)
	time.Sleep(sleepTime)
	fmt.Printf("%d - Finish\n", id)
	
	// Finish
	wg2.Done()
}

const n = 5

func main() {
	var threads [n]Thread
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	for i := 0; i < n; i++ {
		wg1.Add(1)
		wg2.Add(1)
		threads[i] = Thread{i, -1}
		go execute(i, n, &threads, &wg1, &wg2)
	}
	
	wg2.Wait()
	
	fmt.Println("Finish All")
	
}