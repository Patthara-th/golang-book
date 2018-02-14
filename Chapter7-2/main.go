package main

import "fmt"

func main() {
	slice := make([]int, 4, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	fmt.Println("Slice with length and capacity")
	fmt.Printf("slice: length %v, capacity %v, %v\n", len(slice), cap(slice), slice)
	fmt.Printf("slice2: length %v, capacity %v, %v\n", len(slice2), cap(slice2), slice2)

	// append
	for i := 4; i < 15; i++ {
		slice = append(slice, i)
	}

	fmt.Printf("slice: length %v, capacity %v, %v\n", len(slice), cap(slice), slice)

}
