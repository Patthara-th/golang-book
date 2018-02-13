package main

import "fmt"

func main() {
	fmt.Println("Enter Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) / 1.8
	fmt.Printf("Celcius = %.2f", output)
}
