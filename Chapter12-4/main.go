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

	ModFunc := func(a int, c string) func(b int) (string, bool) {
		return func(b int) (string, bool) {
			if b%a == 0 {
				return c, true
			}
			return "", false
		}
	}

	// fbArray := [...]func(n int) (string, bool){
	// 	ModFunc(15, "FizzBuzz"),
	// 	ModFunc(3, "Fizz"),
	// 	ModFunc(5, "Buzz"),
	// }

	getfbArray := func(num [3]int, st [3]string) [3]func(n int) (string, bool) {
		var result = [3]func(n int) (string, bool){
			ModFunc(num[0], st[0]),
			ModFunc(num[1], st[1]),
			ModFunc(num[2], st[2]),
		}
		return result
	}

	ln := [3]int{15, 5, 3}
	str := [3]string{"FizzBuzz", "Buzz", "Fizz"}
	fbArray := getfbArray(ln, str)

	for _, v := range fbArray {
		if str, ok := v(number); ok {
			return str
		}
	}
	return ""

}
