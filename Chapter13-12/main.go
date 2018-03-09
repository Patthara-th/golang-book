package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// go func() {ch <- 4 }()
	// c <- 3 // overfilled the buffer
	fmt.Println(<-c)
	fmt.Println(<-c)

}
