package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func execute(id int, wg *sync.WaitGroup) {
	sleepTime := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(sleepTime)
	fmt.Printf("Thread %d finished\n", id)
	wg.Done()
}

func main() {
	n := 5
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go execute(i, &wg)
	}
	wg.Wait()
	fmt.Println("Threads Finished")
}
