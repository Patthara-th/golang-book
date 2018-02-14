package main

import (
	"fmt"
)

func main() {
	x,y := f1()
	fmt.Println(x,y)
}

func f1() (int, int) {
	return 5,7
}