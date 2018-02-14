package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2, 3))

	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
