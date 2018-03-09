package main

import (
	"fmt"
	// "runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(16)
	for i := 1; i < 17; i++ {
		go increment(i)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(n int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// runtime.Gosched()
		value++
		counter = value

	}
}
