package main

import (
	"fmt"
	"strconv"
)

var s0 = [3]string{" _ ", "| |", "|_|"}
var s1 = [3]string{"   ", "  |", "  |"}
var s2 = [3]string{" _ ", " _|", "|_ "}
var s3 = [3]string{" _ ", " _|", " _|"}
var s4 = [3]string{"   ", "|_|", "  |"}
var s5 = [3]string{" _ ", "|_ ", " _|"}
var s6 = [3]string{"   ", "|_ ", "|_|"}
var s7 = [3]string{" _ ", "  |", "  |"}
var s8 = [3]string{" _ ", "|_|", "|_|"}
var s9 = [3]string{" _ ", "|_|", "  |"}
var sdash = [3]string{"   ", " _ ", "   "}

func main() {
	fmt.Println(weatherCelcius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
	fmt.Println(weatherCelcius(1234567890, "Chiang-mai cold"))
	fmt.Println(weatherCelcius(-81, "Chiang-mai cold"))
}

func weatherCelcius(Temp int, place string) string {

	line := [3]string{"", "", ""}
	var number string = strconv.Itoa(Temp)

	for _, c := range number {

		d1, d2, d3 := getformat(string(c))
		line[0] = line[0] + d1
		line[1] = line[1] + d2
		line[2] = line[2] + d3
	}

	return line[0] + "\n" + line[1] + "\n" + line[2] + " C\n" + place

	// fmt.Println(line[0])
	// fmt.Println(line[1])
	// fmt.Println(line[2] + " C")
	// fmt.Println(place)

}

func getformat(i string) (string, string, string) {
	switch i {
	case "0":
		return s0[0], s0[1], s0[2]
	case "1":
		return s1[0], s1[1], s1[2]
	case "2":
		return s2[0], s2[1], s2[2]
	case "3":
		return s3[0], s3[1], s3[2]
	case "4":
		return s4[0], s4[1], s4[2]
	case "5":
		return s5[0], s5[1], s5[2]
	case "6":
		return s6[0], s6[1], s6[2]
	case "7":
		return s7[0], s7[1], s7[2]
	case "8":
		return s8[0], s8[1], s8[2]
	case "9":
		return s9[0], s9[1], s9[2]
	case "-":
		return sdash[0], sdash[1], sdash[2]
	}
	return "", "", ""
}
