package main

import "fmt"

func main() {
	var x map[string]int
	x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	y := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	y["four"] = 4
	fmt.Println(y)

	z := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println(z)
	fmt.Println(z[1])
}
