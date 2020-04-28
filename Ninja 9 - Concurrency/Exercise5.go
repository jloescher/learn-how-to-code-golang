package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Fix the race condition you created in exercise #4 by using package atomic

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
			atomic.AddInt64(&counter, 1)                           // atomic increment counter by 1
			fmt.Println(i, "Counter:", atomic.LoadInt64(&counter)) // access without race
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
