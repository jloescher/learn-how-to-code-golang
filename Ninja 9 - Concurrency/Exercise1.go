package main

import (
	"fmt"
	"sync"
)

/*
in addition to the main goroutine, launch two additional goroutines
- each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

var wg sync.WaitGroup
var mu sync.Mutex
var counter = 0

func main() {
	wg.Add(3)
	go R1(counter, 100)
	go R2(counter, 100)
	go R3(counter, 100)

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func R1(counter, total int) {
	for counter < total {
		mu.Lock()
		fmt.Println("Go Routine 1:", counter)
		counter++
		mu.Unlock()
	}
	wg.Done()
}

func R2(counter, total int) {
	for counter < total {
		mu.Lock()
		fmt.Println("Go Routine 2:", counter)
		counter++
		mu.Unlock()
	}
	wg.Done()
}

func R3(counter, total int) {
	for counter < total {
		mu.Lock()
		fmt.Println("Go Routine 3:", counter)
		counter++
		mu.Unlock()
	}
	wg.Done()
}
