package main

import (
	"sync/atomic"
	"sync"
	"fmt"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
	mu      sync.Mutex
)


func main() {
	volumn := 500

	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

}

func order(volumn int) (container []string) {
	cashier := make(chan string)
	barista := make(chan string)
	
	baristano := 50

	go func() {
		for x := 1; x <= volumn; x++ {
			cashier <- receiveOrder(x)
		}
		close(cashier)
	}()	
	
	// quit := make(chan int)
	
	
	// go func() {
	// 	for i:= 1 ; i<= volumn; i++ {
	// 		container = append(container,serve(<-barista))
	// 	}
	// 	for i := 1; i <= baristano; i++ { 
	// 		quit <- 0
	// 	}
	// }()

	// wg.Add(baristano)
	// for i := 1; i <= baristano; i++ {
	// 	go conbrew(barista,cashier,quit)
	// }
	// wg.Wait()

	
	for i := 1; i <= baristano; i++ {
		go conbrew1(volumn,barista,cashier)
	}
	
	for i := 1 ; i<= volumn ; i++ {
		container = append(container,serve(<-barista))
	}
	
	return container
}

func receiveOrder(number int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", number)
}

func brew(order string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}

func serve(coffee string) string {
	time.Sleep(5 * time.Millisecond)
	return  fmt.Sprintf("%s %s", coffee, "ready :)")

}

// func conbrew(ch,ch1 chan string, quit chan int) {
	 
// 	for {
// 		select {
// 		case ch <- brew(<-ch1):
			
// 		case <-quit:
// 			wg.Done()
// 			return
// 		}
// 	}
// }

func conbrew1(v int,ch,ch1 chan string) {
	
	x := int64(v)
	for {
		atomic.AddInt64(&counter, 1) 
		if counter <= x { 
			ch <- brew(<-ch1)
		} else {
			
			return 
		} 

	}
	
}


