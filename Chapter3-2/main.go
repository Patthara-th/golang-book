package main

import "fmt"

func main() {
	fmt.Println("=====String=====")
	backticks := `hello wolrd!,
today's good day`

	fmt.Println(backticks)

	doubleQoutes := "hello world!,\ntoday's good day"
	fmt.Println(doubleQoutes)
}
