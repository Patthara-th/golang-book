package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(10)

	//fmt.Printf("Result = %v\n", result)

	count := 0

	for count < 5 {

		fmt.Println("================")
		fmt.Println("Enter Number: ")
		var input int
		fmt.Scanf("%d\n", &input)
		count += 1

		fmt.Printf("Time =  %v\n", count)
		fmt.Printf("Input number = %v", input)

		switch {
		case input == result:
			fmt.Println("  <<เจอแล้ว>>")
			return
		case input < result:
			fmt.Println("  <<น้อยไป>>")
		case input > result:
			fmt.Println("  <<มากไป>>")
		}
	}

	fmt.Println("")
	fmt.Println("*********** เกินพอ ***********")
	fmt.Printf("          Result = %v\n", result)
	fmt.Println("*****************************")
	fmt.Println("")

}
