package main

import (
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(fizzbuzz(number))
	}
}

func fizzbuzz(number int) (int, string) {

	if number%15 == 0 {
		return number, "FizzBuzz"
	} else if number%3 == 0 {
		return number, "Fizz"
	} else if number%5 == 0 {
		return number, "Buzz"
	} else {
		return number, ""
	}

}
