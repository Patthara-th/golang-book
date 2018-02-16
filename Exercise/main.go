package main

import "fmt"

// var s0,s1,s2,s3,s4,s5,s6,s7,s8,s9 [3]string

var s2 = [3]string{" _ "," _|","|_ "}
var s5 = [3]string{" _ ","|_ "," _|"}

func main () {
	weatherCelcius(25,"Bangkok few cloud")

}

func weatherCelcius(Temp int, place string)  {
	

	digit1 := Temp/10
	digit2 := Temp%10

	fmt.Println(digit1)
	fmt.Println(digit2)
	digit11,gigit12,digit13 := getformat(digit1)
	digit21,gigit22,digit23 := getformat(digit2)

	fmt.Println(digit11 , digit21)
	fmt.Println(gigit12 , gigit22)
	fmt.Println(digit13 , digit23 ," C")	
	fmt.Println(place)
	
}

func getformat(i int) (string,string,string) {
	switch i {
	case 2 :
		return s2[0],s2[1],s2[2]
	case 5:
		return s5[0],s5[1],s5[2]
	}

	return "","",""
}