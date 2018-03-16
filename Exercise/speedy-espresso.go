package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
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

	queue := dispatch(volumn)

	cup := pickup(queue)

	return servecoffee(cup, container)

}

func dispatch(volumn int) chan string {
	queue := make(chan string)
	go func() {
		for x := 1; x <= volumn; x++ {
			queue <- receiveOrder(x)
		}
		close(queue)
	}()
	return queue
}

func pickup(queue <-chan string) chan string {
	cup := make(chan string)
	baristano := 20

	wg.Add(baristano)
	for i := 1; i <= baristano; i++ {
		go pullshot(queue, cup)
	}

	go func() {
		wg.Wait()
		close(cup)
	}()

	return cup
}

func pullshot(queue <-chan string, cup chan<- string) {
	for x := range queue {
		cup <- brew(x)
	}
	wg.Done()
}

func servecoffee(cup <-chan string, container []string) []string {

	for x := range cup {
		container = append(container, serve(x))
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
	return fmt.Sprintf("%s %s", coffee, "ready :)")
}
