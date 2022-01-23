package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type Philosopher struct {
	id     int
	status string
}

func getRight(id int) int {
	if (id == 0) {
		return 5 - 1;
	} else {
		return id - 1;
	}
}

func getLeft(id int) int {
	return (id + 1) % 5;
}

func execute(id int, philosophers *[5]Philosopher, 
			 mutex *sync.Mutex, philosophers_sem *[5]sync.Mutex) {
	
	sleepTime := time.Duration(rand.Intn(5)) * time.Second
	
	right := getRight(id)
	left := getLeft(id)
	
	for {
		
		// THINK
		
		fmt.Printf("%d is %s\n", id, philosophers[id].status)
		
		time.Sleep(sleepTime)
		
		// TAKE CLUTERY
		
		mutex.Lock()
		philosophers[id].status = "hungry"
		fmt.Printf("%d is %s\n", id, philosophers[id].status)
		
		if (philosophers[right].status != "eating" && 
				philosophers[left].status != "eating") {
			
			philosophers[id].status = "eating"
			philosophers_sem[id].Unlock()
			
			fmt.Printf("%d is %s - LEFT: %d is %s - Right: %d is %s\n", id, philosophers[id].status, 
					   left, philosophers[left].status,
					   right, philosophers[right].status)
			
		} else {
			fmt.Printf("%d couldn't eat - LEFT: %d is %s - Right: %d is %s\n", id, 
					   left, philosophers[left].status,
					   right, philosophers[right].status)	
		}
		
		mutex.Unlock()
		philosophers_sem[id].Lock()
		
		// EATING
		
		time.Sleep(sleepTime)
		
		// DROP CLUTERY
		
		mutex.Lock()
		
		philosophers[id].status = "thinking"
		
		if (philosophers[right].status == "hungry" && 
				philosophers[getRight(right)].status != "eating") {
			
			philosophers[right].status = "eating"
			philosophers_sem[right].Unlock()
			
			fmt.Printf("%d is %s - LEFT: %d is %s - Right: %d is %s\n", right, philosophers[right].status, 
					   getLeft(right), philosophers[getLeft(right)].status,
					   getRight(right), philosophers[getRight(right)].status)
		}
		
		if (philosophers[left].status == "hungry" && 
				philosophers[getLeft(left)].status != "eating") {
			
			philosophers[left].status = "eating"
			philosophers_sem[left].Unlock()
			
			fmt.Printf("%d is %s - LEFT: %d is %s - Right: %d is %s\n", left, philosophers[left].status, 
					   getLeft(left), philosophers[getLeft(left)].status,
					   getRight(left), philosophers[getRight(left)].status)
			
		}
		
		mutex.Unlock()
	}
}

func main() {
	const n = 5
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var philosophers [n]Philosopher
	var philosophers_sem [n]sync.Mutex
	
	for i := 0; i < n; i++ {
		philosophers[i] = Philosopher{i, "thinking"}
		philosophers_sem[i].Lock()
	}
	
	fmt.Println("Starting...")
	
	for i := 0; i < n; i++ {
		wg.Add(1)
		go execute(i, &philosophers, &mutex, &philosophers_sem)
	}
	
	wg.Wait()
}