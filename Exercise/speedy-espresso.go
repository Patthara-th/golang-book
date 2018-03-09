package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 200

	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

}

func order(volumn int) (container []string) {
	cashier := make(chan string)
	barista := make(chan string)
	
	
	
	go func() {
		for x := 1; x <= volumn; x++ {
			cashier <- receiveOrder(x)
		}
		close(cashier)
	}()
	
	
	go func() {
		for x := range cashier {
			barista <- brew(x)
		}	
		close(barista)	
	}()

	go func() {
		for x := range cashier {
			barista <- brew(x)
		}	
		
	}()

	go func() {
		for x := range cashier {
			barista <- brew(x)
		}	
			
	}()
	
	
	for x := range barista {
			time.Sleep(5* time.Millisecond)
			container = append( container,fmt.Sprintf("%s %s", x, "ready :)"))
	}
		
	



	// for i := 1 ; i <= volumn ; i ++ {

	// 	time.Sleep(5 * time.Millisecond)
	// 	coffee := fmt.Sprintf("order: %d", i)

	// 	time.Sleep(100 * time.Millisecond)
	// 	coffee = fmt.Sprintf("%s %s", coffee, "espresso")

	// 	time.Sleep(5 * time.Millisecond)
	// 	container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))

	// }

	return container}




func receiveOrder(number int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", number)
}

func brew(order string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}

func serve(coffee string) string {
	time.Sleep(5 * time.Millisecond)
	return  fmt.Sprintf("%s %s", coffee, "ready :)")

}


