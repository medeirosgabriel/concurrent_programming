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

func execute(id int, n int, threads *[5]Thread, sems *[5]sync.Mutex, wg *sync.WaitGroup) {

	// Sleeping
	time_seconds := rand.Intn(5)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds\n", id, time_seconds)
	time.Sleep(sleepTime)

	fmt.Printf("%d - Waking Up\n", id)

	// Choosing next time
	time_seconds = rand.Intn(5)
	threads[id].nextTime = time_seconds
	fmt.Printf("%d - Next Time = %d\n", id, time_seconds)
	nextId := getNextId(id, n)
	sems[nextId].Unlock()
	sems[id].Lock()

	// Sleeping Again
	lastId := getLastId(id, n)
	time_seconds = threads[lastId].nextTime
	sleepTime = time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds - Second Time\n", id, time_seconds)
	time.Sleep(sleepTime)

	// Done
	wg.Done()
}

func main() {
	const n = 5
	var threads [n]Thread
	var sems [n]sync.Mutex
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		threads[i] = Thread{i, -1}
		sems[i].Lock()
		go execute(i, n, &threads, &sems, &wg)
	}

	wg.Wait()
}
