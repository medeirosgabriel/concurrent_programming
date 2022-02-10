package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func exec(id int, sem chan int, wg *sync.WaitGroup) {
	time_seconds := rand.Intn(10)
	sleepTime := time.Duration(time_seconds) * time.Second
	fmt.Printf("%d - Sleeping for %d seconds\n", id, time_seconds)
	time.Sleep(sleepTime)
	fmt.Printf("%d - Finished\n", id)
	<-sem
	wg.Done()
}

func main() {
	n := 5
	var wg sync.WaitGroup
	sem := make(chan int, n)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		sem <- 1
		go exec(i, sem, &wg)
	}
	fmt.Println("Waiting Threads!!")
	wg.Wait()
	fmt.Println("=== FINISH ALL ===")

}
