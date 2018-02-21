package main

import (
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
	}
}

func fizzbuzz(number int) string {

	ln := [3]int{15, 3, 5}
	str := [3]string{"FizzBuzz", "Fizz", "Buzz"}

	for i := 0; i < 3; i++ {
		if number%ln[i] == 0 {
			return str[i]
		}

	}
	// if number%ln[0] == 0 {
	// 	return str[0]
	// }

	// if number%ln[1] == 0 {
	// 	return str[1]
	// }

	// if number%ln[2] == 0 {
	// 	return str[2]
	// }

	return ""

}
